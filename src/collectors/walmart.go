package collectors

import (
	"fmt"
	"strings"

	"github.com/ManoloTonto1/Poopie-Hadoop/hadoop"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func ScrapeWalmart() {
	product := hadoop.Product{}
	mainCollector := colly.NewCollector(
		colly.AllowedDomains("www.walmart.com"),
		colly.UserAgent(userAgent),
	)
	mainCollector.Limit(limitRules)
	productCollector := mainCollector.Clone()

	mainCollector.OnHTML("a.absolute.w-100.h-100.z-1.hide-sibling-opacity", func(e *colly.HTMLElement) {
		_link := e.Attr("href")
		productCollector.Visit(e.Request.AbsoluteURL(_link))

	})

	productCollector.OnHTML("h1.b.lh-copy.dark-gray.mt1.mb2.f3", func(e *colly.HTMLElement) {

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

	productCollector.OnHTML("body", func(h *colly.HTMLElement) {
		image := h.DOM.Find("img.noselect.db").First().AttrOr("src", "")
		fmt.Println("image: " + image)
		product.Image = image

		h.DOM.Find("div.w_HmLO").Each(func(i int, s *goquery.Selection) {
			text := s.Find("span.tl-m.db-m > span").Text()
			if text != "" && len(strings.Split(text, " ")) >= 100 {
				product.Reviews = append(product.Reviews, text)
			}
		})
		if product.Reviews == nil {
			return
		}
		// 	if err := hadoop.CreateFile("products", product); err != nil {
		// 		panic(err)
		// 	}
		// product = hadoop.Product{}

	})
	mainCollector.OnHTML("a.b--light-gray", func(e *colly.HTMLElement) {
		_link := e.Attr("href")
		productCollector.Visit(e.Request.AbsoluteURL(_link))

	})
	mainCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	mainCollector.Visit("https://www.walmart.com/search?q=outdoor+shoes&typeahead=outdoor+shoes&catId=5438")

}
