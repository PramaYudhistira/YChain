package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		fmt.Println("Title:", e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "Error:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting ==============================", r.URL)
	})

	err := c.Visit("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal("Failed to visit website:", err)
	}

}
