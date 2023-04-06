package collectors

import (
	"fmt"
	"strings"
	"time"

	"github.com/ManoloTonto1/Poopie-Hadoop/hadoop"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func Init() {
	var products = []hadoop.Product{}

	c := colly.NewCollector(
		colly.Async(true),
		colly.AllowedDomains("www.amazon.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"),
	)
	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
		Parallelism: 4,
	})

	c.OnHTML("a.a-link-normal.s-no-outline", func(e *colly.HTMLElement) {
		product := hadoop.Product{}
		_link := e.Attr("href")
		e.Request.Visit(_link)

		title := e.DOM.Find("span#productTitle").First().Text()
		if product.Title != "" {
			return
		}
		fmt.Println("title: " + title)
		product.Title = strings.TrimLeft(title, " ")
		product.Title = strings.TrimRight(product.Title, " ")
		product.Title = strings.ReplaceAll(product.Title, "'", "")

		image := e.DOM.Find("div#unrolledImgNo0 >div.a-image-container > img").First().AttrOr("src", "")
		if product.Image != "" {
			return
		}
		product.Image = image

		link := e.DOM.Find("a.a-link-emphasis.a-text-bold").AttrOr("href", "")
		e.Request.Visit(link)

		e.DOM.Find("span.a-size-base.review-text.review-text-content > span").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if text == "" || len(text) < 100 {
				return
			}
			product.Reviews = append(product.Reviews, text)
		})
		products = append(products, product)

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		if len(product.Reviews) == 0 {
			return
		}
		if err := hadoop.CreateFile("products", product); err != nil {
			panic(err)
		}
		product = hadoop.Product{}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://www.amazon.com/s?k=outdoor+shoes&s=review-rank&ds=v1%3A5iK8eNCtJ3%2FFRhDW%2FZjDHZOAB7AdjYiLsaY513OaKEo&crid=3H44VUYJUM6FI&qid=1680523185&sprefix=outdoor+shoes%2Caps%2C321&ref=sr_st_review-rank")
	c.Wait()
}

// c.OnHTML("li.a-last > a", func(e *colly.HTMLElement) {
// 	link := e.Attr("href")
// 	c.Visit(link)
// })
// }
