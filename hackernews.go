package main

import (
	"encoding/json"
	"os"

	"github.com/gocolly/colly"
)

type Post struct {
	Title string `json:"title"`
	Site string `json:"site"`
}

func main() {
	posts := make([]*Post, 0)

	// Instantiate default collector
	c := colly.NewCollector()

	// Extract post
	c.OnHTML("table.itemlist tbody tr.athing", func(e *colly.HTMLElement) {
		post := &Post{
			Title:e.ChildText("td.title a"),
			Site:e.ChildText("td.title span a"),
		}
		posts = append(posts, post)
	})

	c.Visit("https://news.ycombinator.com/")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(posts)
}