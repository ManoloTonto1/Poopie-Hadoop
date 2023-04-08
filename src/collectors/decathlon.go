package collectors

import (
	"fmt"
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
	mainCollector.Cookies(decathlonCookieString)
	mainCollector.Limit(limitRules)
	productCollector := mainCollector.Clone()
	reviewCollector := mainCollector.Clone()

	mainCollector.OnHTML("a.dpb-product-model-link", func(e *colly.HTMLElement) {
		_link := e.Attr("href")
		productCollector.Visit(e.Request.AbsoluteURL(_link))

	})

	productCollector.OnHTML("body", func(e *colly.HTMLElement) {

		product.URl = e.Request.AbsoluteURL(e.Request.URL.String())
		fmt.Println("link: " + product.URl)
		title := e.DOM.Find("h1.vtmn-typo_title-4").Text()
		fmt.Println("title: " + title)
		if product.Title != "" {
			return
		}
		product.Title = strings.TrimLeft(title, " ")
		product.Title = strings.TrimRight(product.Title, " ")
		product.Title = strings.ReplaceAll(product.Title, "'", "")
		product.Title = strings.ReplaceAll(product.Title, "/", "_")

		image := e.DOM.Find("section.active:nth-child(2) > img").AttrOr("src", "")
		fmt.Println("image: " + image)
		product.Image = image
	})

	productCollector.OnHTML("a.cta:nth-child(5)", func(h *colly.HTMLElement) {
		reviewCollector.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})

	reviewCollector.OnHTML("body", func(h *colly.HTMLElement) {
		h.Request.Headers = &decathlonHeaders
		h.Request.Do()
			h.DOM.Find("article.review-article-container > div > em").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			fmt.Println("text: " + text)
			if text != "" && len(strings.Split(text, " ")) >= 100 {
				product.Reviews = append(product.Reviews, text)
			}
		})
		if product.Reviews == nil {
			return
		}
		// if err := hadoop.CreateFile("products", product); err != nil {
		// 	panic(err)
		// }
		// product = hadoop.Product{}
	})

	mainCollector.OnHTML("a.s-pagination-item:nth-child(8)", func(e *colly.HTMLElement) {
		_link := e.Attr("href")
		productCollector.Visit(e.Request.AbsoluteURL(_link))

	})
	mainCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	mainCollector.Visit("https://www.decathlon.nl/search?Ntt=schoenen&facets=sportGroupLabels:Hiking_natureLabel:Schoenen_")

}
