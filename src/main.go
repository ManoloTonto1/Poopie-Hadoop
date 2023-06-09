package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/ManoloTonto1/Poopie-Hadoop/analysis"
	"github.com/ManoloTonto1/Poopie-Hadoop/collectors"
	"github.com/ManoloTonto1/Poopie-Hadoop/hadoop"
)

type Logs struct {
	Products uint
	Reviews  uint
}

func LogData(startTime time.Time) {
	logs := Logs{}
	client, err := hadoop.InitConnectionWithHDFSCluster()
	if err != nil {
		fmt.Println("Error initializing connection with hdfs cluster: ", err)
	}
	defer client.Close()
	products, err := client.ReadDir("/products")
	if err != nil {
		fmt.Println("Error reading products directory: ", err)
		return
	}
	logs.Products = uint(len(products))
	for _, product := range products {
		jsonData := hadoop.Product{}
		data, err := client.ReadFile("/products/" + product.Name())
		if err != nil {
			fmt.Println("Error reading product file: ", err)
			return
		}
		if json.Unmarshal(data, &jsonData); err != nil {
			fmt.Println("Error unmarshaling product file: ", err)
			return
		}
		logs.Reviews += uint(len(jsonData.Reviews))
	}
	fmt.Println("Time Taken: ", time.Since(startTime))
	fmt.Println("Products: ", logs.Products)
	fmt.Println("Reviews: ", logs.Reviews)
	fmt.Println("All Jobs Done! Closing Connections")
}

func main() {
	wg := sync.WaitGroup{}
	startTime := time.Now()
	wg.Add(3)
	go func() {
		defer wg.Done()
		collectors.ScrapeAmazon()
	}()
	go func() {
		defer wg.Done()
		collectors.ScrapeDecathlon()
	}()
	go func() {
		defer wg.Done()
		collectors.ScrapeWalmart()
	}()
	wg.Wait()
	LogData(startTime)
	analysis.LoadEnv()
	analysis.ConnectToDB()
	analysis.InitAnalysis()
	analysis.StartDashBoard()

}
