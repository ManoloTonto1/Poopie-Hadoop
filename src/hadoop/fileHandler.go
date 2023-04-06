package hadoop

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/colinmarc/hdfs"
)

var nameNode = "localhost:9000"

func InitConnectionWithHDFSCluster() (*hdfs.Client, error) {
	client, err := hdfs.NewClient(
		hdfs.ClientOptions{Addresses: []string{
			nameNode,
		}})
	if err != nil {
		return nil, fmt.Errorf("error creating hdfs client: %w", err)
	}
	return client, nil
}

func HandlePathResolution(client *hdfs.Client, directory string, fileName string) (string, error) {
	rootFileInfo, err := client.Stat("/")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	rootAbsolutePath := rootFileInfo.Name()

	dirPath := rootAbsolutePath + directory
	if err := client.MkdirAll(dirPath, 0777); err != nil {
		return "", fmt.Errorf("error creating directory %s: %w", directory, err)
	}
	completePath := dirPath + "/" + strings.ReplaceAll(fileName, " ", "_")
	return completePath, nil
}

func CreateFile(directory string, data Product) error {
	client, err := InitConnectionWithHDFSCluster()

	if err != nil {
		return fmt.Errorf("error initializing connection with hdfs cluster: %w", err)
	}
	defer client.Close()

	fileContent, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling file content: %w", err)
	}

	filePath, err := HandlePathResolution(client, directory, data.Title)
	if err != nil {
		return fmt.Errorf("error resolving file path: %w", err)
	}

	filePath = filePath + ".json"
	if file, err := client.Stat(filePath); file == nil && err == nil {
		return nil
	}
	file, err := client.CreateFile(filePath, 1, 1048576, 0777)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filePath, err)
	}
	defer file.Close()

	if _, err = file.Write(fileContent); err != nil {
		return fmt.Errorf("error writing to file %s: %w", filePath, err)
	}

	return nil
}
