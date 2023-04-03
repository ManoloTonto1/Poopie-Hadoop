package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Product struct {
	Title   string
	Image   string
	Reviews []string
}

func (p *Product) GetReviews(c *colly.Collector) {
	c.OnHTML("span.a-size-base.review-text.review-text-content > span", func(e *colly.HTMLElement) {
		text := e.Text
		if len(text) < 500 {
			return
		}
		// fmt.Println("text: " + text)
		p.Reviews = append(p.Reviews, text)
	})
	c.OnHTML("li.a-last > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(link)
		p.GetReviews(c)
	})
}

func GetDetailsFromPage(c *colly.Collector, link string) {
	product := Product{}
	c.Visit(link)
	c.OnHTML("span#productTitle", func(e *colly.HTMLElement) {
		title := e.Text
		if product.Title != "" {
			return
		}
		fmt.Println("title: " + title)
		product.Title = title

	})

	c.OnHTML("div.a-image-container > img", func(e *colly.HTMLElement) {
		image := e.Attr("src")
		if product.Image != "" {
			return
		}
		// fmt.Println("image: " + image)
		product.Image = image
	})

	// travel to reviews page
	c.OnHTML("a.a-link-emphasis.a-text-bold", func(e *colly.HTMLElement) {
		if e.Text == "See all reviews" {
			link := e.Attr("href")
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	product.GetReviews(c)
    if len(product.Reviews) == 0 {
        return
    }

	file, _ := json.MarshalIndent(product, "", " ")
	_ = ioutil.WriteFile("products/"+product.Title, file, 0644)

}

func main() {
	// Create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"),
	)
	pageCollector := c.Clone()

	c.OnHTML("html", func(e *colly.HTMLElement) {
		e.ForEach("a.a-link-normal.s-no-outline", func(_ int, e *colly.HTMLElement) {
			link := e.Attr("href")
			GetDetailsFromPage(pageCollector, e.Request.AbsoluteURL(link))
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://www.amazon.com/s?k=outdoor+shoes&s=review-rank&ds=v1%3A5iK8eNCtJ3%2FFRhDW%2FZjDHZOAB7AdjYiLsaY513OaKEo&crid=3H44VUYJUM6FI&qid=1680523185&sprefix=outdoor+shoes%2Caps%2C321&ref=sr_st_review-rank")
}
