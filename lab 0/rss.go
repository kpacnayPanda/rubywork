package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func main() {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://news.mail.ru/rss/90/")
	fmt.Println("Title: ", feed.Title)
	fmt.Println("Generator: ", feed.Generator)
	fmt.Println("Updated: ", feed.Updated)
	fmt.Println("Description: ", feed.Description)
	//fmt.Println("\nNumber of Items:", feed.Items)

	for v := range feed.Items {
		item := feed.Items[v]
		fmt.Println();
		fmt.Println("Item Number", v)
		fmt.Println("Title:", item.Title)
		fmt.Println("Link:", item.Link)
		fmt.Println("Description:", item.Description)
	}
}