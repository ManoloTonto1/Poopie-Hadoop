package collectors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ManoloTonto1/Poopie-Hadoop/hadoop"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func ScrapeDecathlon() {
	product := hadoop.Product{}
	mainCollector := colly.NewCollector(
		colly.AllowedDomains("www.decathlon.nl"),
		colly.UserAgent(userAgent),
	)
	mainCollector.Limit(limitRules)
	productCollector := mainCollector.Clone()

	mainCollector.OnHTML("body", func(e *colly.HTMLElement) {
		e.DOM.Find("a.dpb-product-model-link").Each(func(i int, s *goquery.Selection) {
			_link := s.AttrOr("href", "")
			productCollector.Visit(e.Request.AbsoluteURL(_link))
		})

		nextPage := e.DOM.Find("a.s-pagination-item:nth-child(8)").AttrOr("href", "")
		productCollector.Visit(e.Request.AbsoluteURL(nextPage))
	})

	productCollector.OnHTML("body", func(e *colly.HTMLElement) {

		title := e.DOM.Find("h1.vtmn-typo_title-4").Text()
		fmt.Println("title: " + title)
		product.Title = strings.TrimLeft(title, " ")
		product.Title = strings.TrimRight(product.Title, " ")
		product.Title = strings.ReplaceAll(product.Title, "'", "")
		product.Title = strings.ReplaceAll(product.Title, "/", "_")

		image := e.DOM.Find("section.active:nth-child(2) > img").AttrOr("src", "")
		fmt.Println("image: " + image)
		product.Image = image
		reviewsId := e.DOM.Find(`div[id^="ProductReviews-"]`).First().AttrOr("id", "not found")
		fmt.Println("reviewsId: " + reviewsId)
		product.URl = e.Request.AbsoluteURL(e.Request.URL.String())
		fmt.Println("link: " + product.URl)

		body := DecathlonPostData{
			Components: []struct {
				ID    string "json:\"id\""
				Input struct {
					AsyncRequest bool     "json:\"asyncRequest\""
					Count        int      "json:\"count\""
					Ids          []string "json:\"ids\""
					Page         int      "json:\"page\""
				} "json:\"input\""
				Type string "json:\"type\""
			}{
				{
					ID: reviewsId,
					Input: struct {
						AsyncRequest bool     "json:\"asyncRequest\""
						Count        int      "json:\"count\""
						Ids          []string "json:\"ids\""
						Page         int      "json:\"page\""
					}{
						AsyncRequest: true,
						Count:        1000,
						Ids:          []string{strings.Split(product.URl, "?mc=")[1]},
						Page:         1,
					},
					Type: "ProductReviews",
				},
			},
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", "https://www.decathlon.nl/nl/ajax/nfs/async", bytes.NewBuffer(jsonBody))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		// Send the request and get the response
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Read the response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		responseStatus := resp.StatusCode
		fmt.Println(responseStatus)
		if responseStatus == 200 {
			responseJson := []DecathlonResponseBody{}
			responseBody := responseBody
			if err := json.Unmarshal(responseBody, &responseJson); err != nil {
				panic(err)
			}
			reviews := responseJson[0].Num0.Data.Reviews
			for _, review := range reviews {
				fmt.Println("review: " + review.Review.Body)
				product.Reviews = append(product.Reviews, review.Review.Body)
			}
			if err := hadoop.CreateFile("products", product); err != nil {
				panic(err)
			}

		}
	})

	mainCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	mainCollector.Visit("https://www.decathlon.nl/search?Ntt=schoenen&facets=sportGroupLabels:Hiking_natureLabel:Schoenen_")
}
