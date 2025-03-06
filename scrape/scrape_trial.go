package scrape

import (
	"github.com/gocolly/colly"
)

func ScrapeTrial() {
	c := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)

}
