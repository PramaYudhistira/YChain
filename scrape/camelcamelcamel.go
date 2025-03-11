package scrape

import (
	"github.com/gocolly/colly"
)

type CamelScraper struct {
	Collector *colly.Collector
}

func NewCamelScraper() *CamelScraper {
	c := colly.NewCollector()
	return &CamelScraper{
		Collector: c,
	}
}

func (cs *CamelScraper) Scrape(url string) {

}
