package main

import (
	"fmt" // пакет для форматированного ввода вывода
	"net/http" // пакет для поддержки HTTP протокола
	"github.com/mmcdole/gofeed" 
	"log"
) 

func MailNews(w http.ResponseWriter, r *http.Request) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://news.mail.ru/rss/90/")

   	//вывод списка ссылок
  	if err == nil {
   		fmt.Fprintf(w, "<html>")
   		fmt.Fprintf(w, "<body>")
		for v := range feed.Items {
			item := feed.Items[v]
			fmt.Println("Num:", v)
			fmt.Println("Title:", item.Title)
			fmt.Fprintf(w, "<b align=center><font size = \"2\"> %s </font></b><br>", item.Title)
			fmt.Fprintf(w, "%s<br>", item.Description)
			fmt.Fprintf(w, "<a href = \"%s\"> Подробнее </a><br>", item.Link) 
			fmt.Println("Link:", item.Link)
			fmt.Fprintln(w, "<br />") 
 		}
		fmt.Fprintf(w, "</body>")
		fmt.Fprintf(w, "</html>")
	}
}

func LentaNews(w http.ResponseWriter, r *http.Request) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://lenta.ru/rss")

   	//вывод списка ссылок
  	if err == nil {
   		fmt.Fprintf(w, "<html>")
   		fmt.Fprintf(w, "<body>")
		for v := range feed.Items {
			item := feed.Items[v]
			fmt.Println("Num:", v)
			fmt.Println("Title:", item.Title)
			fmt.Fprintf(w, "<b align=center><font size = \"2\"> %s </font></b><br>", item.Title)
			fmt.Fprintf(w, "%s<br>", item.Description)
			fmt.Fprintf(w, "<a href = \"%s\"> Подробнее </a><br>", item.Link) 
			fmt.Println("Link:", item.Link)
			fmt.Fprintln(w, "<br />") 
 		}
		fmt.Fprintf(w, "</body>")
		fmt.Fprintf(w, "</html>")
	}
}

func Vzru(w http.ResponseWriter, r *http.Request) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://vz.ru/rss.xml")

   	//вывод списка ссылок
  	if err == nil {
   		fmt.Fprintf(w, "<html>")
   		fmt.Fprintf(w, "<body>")
		for v := range feed.Items {
			item := feed.Items[v]
			fmt.Println("Num:", v)
			fmt.Println("Title:", item.Title)
			fmt.Fprintf(w, "<b align=center><font size = \"2\"> %s </font></b><br>", item.Title)
			fmt.Fprintf(w, "%s<br>", item.Description)
			fmt.Fprintf(w, "<a href = \"%s\"> Подробнее </a><br>", item.Link) 
			fmt.Println("Link:", item.Link)
			fmt.Fprintln(w, "<br />") 
 		}
		fmt.Fprintf(w, "</body>")
		fmt.Fprintf(w, "</html>")
	}
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"http://localhost:4000/mail\"><font size = \"4\"> Новости mail.ru </font></a><br><br>")
	fmt.Fprintf(w, "<a href=\"http://localhost:4000/lenta\"><font size = \"4\"> Новости LentaNews </font></a><br><br>")
	fmt.Fprintf(w, "<a href=\"http://localhost:4000/vz\"><font size = \"4\"> RSS канал vz.ru </font></a><br><br>")
		//вывод общего списка
	
}

func main() {
	http.HandleFunc("/", HomeRouterHandler) // установим роутер
	http.HandleFunc("/mail", MailNews)
	http.HandleFunc("/lenta", LentaNews)
	http.HandleFunc("/vz", Vzru)
	err := http.ListenAndServe(":4000", nil) // задаем слушать порт
	if err != nil { 
		log.Fatal("ListenAndServe: ", err)
	}
}