package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

type Item struct {
	Link    string `json:"link"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Instock string `json:"instock"`
}

func main() {
	defer timer("main")()
	// Create a new Colly collector
	c := colly.NewCollector(colly.Async(true))

	items := []Item{}

	c.OnHTML("div.side_categories li ul li", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		c.Visit(h.Request.AbsoluteURL(link))
	})

	// Define a callback for when an HTML element is found
	// c.OnHTML("li.next a", func(h *colly.HTMLElement) {
	// 	c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	// })

	c.OnHTML("article.product_pod", func(h *colly.HTMLElement) {
		i := Item{
			Link:    h.ChildAttr("a", "href"),
			Name:    h.ChildAttr("h3 a", "title"),
			Price:   h.ChildText("p.price_color"),
			Instock: h.ChildText("p.instock"),
		}
		items = append(items, i)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Start the scraping process
	err := c.Visit("https://books.toscrape.com/catalogue/category/books_1/index.html")
	if err != nil {
		log.Fatal(err)
	}
	c.Wait()

	data, err := json.MarshalIndent(items, " ", "")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./books.json", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s Scrapping took %v", name, time.Since(start))
	}
}
