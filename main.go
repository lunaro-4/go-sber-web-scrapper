package main

import (

	"fmt"

	"github.com/chromedp/chromedp"

	"github.com/gocolly/colly"

)



func main()  {
	chromedp.Reload()
	fmt.Print( colly.ProxyURLKey)
	coll := colly.NewCollector()

	coll.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error: %s\n", err.Error())})

	coll.OnHTML("title", func(h *colly.HTMLElement) {
		fmt.Print(h.Text)
	})
	
	coll.Visit("https://books.toscrape.com/")
}

