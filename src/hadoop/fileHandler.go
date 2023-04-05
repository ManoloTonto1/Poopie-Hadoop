package hadoop

import (
	"encoding/json"
	"fmt"

	"github.com/colinmarc/hdfs"
)

func InitConnectionWithHDFSCluster() (*hdfs.Client, error) {
	client, err := hdfs.New("localhost:9000")
	if err != nil {
		return nil, fmt.Errorf("error creating hdfs client: %w", err)
	}
	return client, nil
}
func HandlePathResolution(client *hdfs.Client, directory string, fileName string) (string, error) {
	if _, err := client.Stat(directory); err != nil {
		if err := client.MkdirAll(directory, 0777); err != nil {
			return "", fmt.Errorf("error creating directory %s: %w", directory, err)
		}
	}
	return fmt.Sprintf("%s/%s", directory, fileName), nil
}

func CreateFile(directory string, data Product) error {
	client, err := InitConnectionWithHDFSCluster()

	if err != nil {
		return fmt.Errorf("error initializing connection with hdfs cluster: %w", err)
	}
	fileContent, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("error marshaling file content: %w", err)
	}

	filePath, err := HandlePathResolution(client, directory, data.Title)
	if err != nil {
		return fmt.Errorf("error resolving file path: %w", err)
	}

	file, err := client.CreateFile(filePath, 2, 1, 0777)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filePath, err)
	}
	defer file.Close()

	if _, err := file.Write(fileContent); err != nil {
		return fmt.Errorf("error writing file content to %s: %w", filePath, err)
	}
	client.Close()
	return nil
}
