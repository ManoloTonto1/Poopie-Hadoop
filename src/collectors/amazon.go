package collectors

import (
	"fmt"
	"strings"
	"time"

	"github.com/ManoloTonto1/Poopie-Hadoop/hadoop"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func FindByLink(products []hadoop.Product, link string) *hadoop.Product {
	for i := 0; i < len(products); i++ {
		if products[i].URl == link {
			return &products[i]
		}
	}
	panic("Product not found")
}

func Init() {
	product := hadoop.Product{}
	mainCollector := colly.NewCollector(
		colly.AllowedDomains("www.amazon.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"),
	)
	mainCollector.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
		Parallelism: 4,
	})
	productCollector := mainCollector.Clone()
	reviewCollector := mainCollector.Clone()

	mainCollector.OnHTML("a.a-link-normal.s-no-outline", func(e *colly.HTMLElement) {
		_link := e.Attr("href")
		productCollector.Visit(e.Request.AbsoluteURL(_link))

	})

	productCollector.OnHTML("span#productTitle", func(e *colly.HTMLElement) {

		product.URl = e.Request.AbsoluteURL(e.Request.URL.String())
		fmt.Println("link: " + product.URl)
		title := e.Text
		fmt.Println("title: " + title)
		if product.Title != "" {
			return
		}
		product.Title = strings.TrimLeft(title, " ")
		product.Title = strings.TrimRight(product.Title, " ")
		product.Title = strings.ReplaceAll(product.Title, "'", "")
		product.Title = strings.ReplaceAll(product.Title, "/", "_")
	})

	productCollector.OnHTML("#unrolledImgNo0 > div:nth-child(1) > img:nth-child(1)", func(h *colly.HTMLElement) {
		image := h.Attr("src")
		fmt.Println("image: " + image)
		product.Image = image
	})

	productCollector.OnHTML("#reviews-medley-footer > div:nth-child(2) > a", func(h *colly.HTMLElement) {
		reviewCollector.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})

	reviewCollector.OnHTML("body", func(h *colly.HTMLElement) {
		h.DOM.Find("span.a-size-base.review-text.review-text-content > span").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if text != "" && len(strings.Split(text, " ")) >= 100 {
				product.Reviews = append(product.Reviews, text)
			}
		})
		if product.Reviews == nil {
			return
		}
		if err := hadoop.CreateFile("products", product); err != nil {
			panic(err)
		}
		product = hadoop.Product{}
	})

	mainCollector.OnHTML("a.s-pagination-item:nth-child(8)", func(e *colly.HTMLElement) {
		_link := e.Attr("href")
		productCollector.Visit(e.Request.AbsoluteURL(_link))

	})
	mainCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	mainCollector.Visit("https://www.amazon.com/s?k=outdoor+shoes&s=review-rank&ds=v1%3A5iK8eNCtJ3%2FFRhDW%2FZjDHZOAB7AdjYiLsaY513OaKEo&crid=3H44VUYJUM6FI&qid=1680523185&sprefix=outdoor+shoes%2Caps%2C321&ref=sr_st_review-rank")

}

// c.OnHTML("li.a-last > a", func(e *colly.HTMLElement) {
// 	link := e.Attr("href")
// 	c.Visit(link)
// })
// }
