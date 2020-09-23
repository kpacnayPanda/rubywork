package main

import (
	"fmt" // пакет для форматированного ввода вывода
	"net/http" // пакет для поддержки HTTP протокола
	"github.com/mmcdole/gofeed" 
	"log"
	"sort"
) 


func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fp1 := gofeed.NewParser()
	feed1, err1 := fp1.ParseURL("https://news.mail.ru/rss/90/")

	fp2 := gofeed.NewParser()
	feed2, err2 := fp2.ParseURL("https://lenta.ru/rss")

	var array[] gofeed.Item

	for i := range feed1.Items {
		array = append(array, *feed1.Items[i])
	}


	for i := range feed2.Items {
		array = append(array, *feed2.Items[i])
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i].Published < array[j].Published
	})
   	//вывод списка ссылок
  	if (err1 == nil) && (err2 == nil) {
   		fmt.Fprintf(w, "<html>")
   		fmt.Fprintf(w, "<body>")
		for v := range array {
			item := array[v]
			fmt.Println("Num:", v)
			fmt.Println("Title:", item.Title)
			fmt.Fprintf(w, "<b align=center><font size = \"2\"> %s </font></b><br>", item.Title)
			fmt.Fprintf(w, "%s<br>", item.Description)
			fmt.Fprintf(w, "%s<br>", item.Published)
			fmt.Fprintf(w, "<a href = \"%s\"> Подробнее </a><br>", item.Link) 
			fmt.Println("Link:", item.Link)
			fmt.Fprintln(w, "<br />") 
 		}
 		fmt.Fprintf(w, "</body>")
		fmt.Fprintf(w, "</html>")
	}
}

func main() {
	http.HandleFunc("/", HomeRouterHandler) // установим роутер
	err := http.ListenAndServe(":4000", nil) // задаем слушать порт
	if err != nil { 
		log.Fatal("ListenAndServe: ", err)
	}
}