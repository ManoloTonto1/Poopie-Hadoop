package analysis

import (
	"fmt"
	"net/http"
	"text/template"
)

type ChartData struct {
	Labels []string
	Data   []int
}

func StartDashBoard() {
	// Create a new instance of the Data struct with some sample data
	reviews := []Review{}
	if err := db.Model(&Review{}).Find(&reviews).Error; err != nil {
		panic(err)
	}
	images := []Image{}
	if err := db.Model(&Image{}).Find(&images).Error; err != nil {
		panic(err)
	}

	// Extract the data for the chart
	labels := make([]string, len(reviews))
	data := make([]int, len(reviews))
	for i, review := range reviews {
		labels[i] = review.Sentiment
		data[i] = review.Score
	}

	chartData := ChartData{
		Labels: labels,
		Data:   data,
	}

	tmpl, err := template.ParseFiles("./dashboard.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Define a handler function for the root URL path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Execute the template with the chart data
		err := tmpl.Execute(w, chartData)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	})

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
