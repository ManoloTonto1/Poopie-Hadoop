package main

import (
	"fmt"
	"time"

	"github.com/ManoloTonto1/Poopie-Hadoop/collectors"
)

type Logs struct {
	Products uint
	Reviews  uint
}

func LogData(logs *Logs, startTime time.Time) {
	// print time taken
	fmt.Println("Time Taken: ", time.Since(startTime))
	fmt.Println("Products: ", logs.Products)
	fmt.Println("Reviews: ", logs.Reviews)
	fmt.Println("All Jobs Done! Closing Connections")
}
func main() {
	startTime := time.Now()
	Logs := Logs{}
	collectors.Init()
	LogData(&Logs, startTime)
}
