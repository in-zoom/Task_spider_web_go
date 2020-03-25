package resources

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func Reading(scheme, domainsName string) {
	resultUrl := scheme + "://" + domainsName + "/"

	c := colly.NewCollector(
		colly.AllowedDomains(domainsName),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Найдена ссылка: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Limit(&colly.LimitRule{
		DomainGlob: domainsName,
		Delay:      1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Посещение", r.URL.String())
	})

	c.Visit(resultUrl)
}
