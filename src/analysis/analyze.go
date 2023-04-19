package analysis

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strings"
	"sync"

	"github.com/EdlinOrg/prominentcolor"
	"github.com/ManoloTonto1/Poopie-Hadoop/hadoop"
)

var wg sync.WaitGroup

func AnalyzeImages(imageUrl string) {
	if imageUrl == "" {
		return
	}
	req, err := http.NewRequest("GET", imageUrl, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting image: ", err)
		return
	}
	defer res.Body.Close()
	// Create a new image from the response body
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error decoding image: ", err)
		return
	}

	prominentColor, err := prominentcolor.Kmeans(img)
	if err != nil {
		fmt.Println("Error getting prominent color: ", err)
		return
	}
	for _, color := range prominentColor {
		if err := db.Create(&Image{
			ImageURL:       imageUrl,
			ProminentColor: CheckRangesForColors(color.AsString()),
		}).Error; err != nil {
			panic(err)
		}
	}
}
func AnalyzeReviews(Reviews []string) {
	positiveWords := append(PositiveWordsDutch, PositiveWordsEnglish...)
	NegativeWords := append(NegativeWordsDutch, NegativeWordsEnglish...)
	for _, review := range Reviews {
		sentiment := "Neutral"
		score := 0
		if len(review) == 0 || review == "" || review == " " || review == "\n" {
			continue
		}

		for _, negativeWord := range NegativeWords {
			if strings.Contains(review, negativeWord) {
				score--
			}
		}
		for _, positiveWord := range positiveWords {
			if strings.Contains(review, positiveWord) {
				score++
			}

		}

		if score > 0 {
			sentiment = "Positive"
		}
		if score < 0 {
			sentiment = "Negative"
		}
		if err := db.Create(&Review{
			Review:    review,
			Score:     score,
			Sentiment: sentiment,
			WordCount: len(review),
		}).Error; err != nil {
			panic(err)
		}
	}

}

func InitAnalysis() {
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
	fmt.Println("Products: ", len(products))
	for _, product := range products {
		wg.Add(2)
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
		go func() {
			defer wg.Done()
			AnalyzeReviews(jsonData.Reviews)
		}()
		go func() {
			defer wg.Done()
			AnalyzeImages(jsonData.Image)
		}()
	}
	wg.Wait()
}
