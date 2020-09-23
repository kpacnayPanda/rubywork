package main

import (
	"fmt" // пакет для форматированного ввода вывода
	"net/http" // пакет для поддержки HTTP протокола
	"strings" // пакет для работы с UTF-8 строками
	"log"// пакет для логирования
	"io"
	"os"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,
	fmt.Println(r.Form) // ввод информации о форме на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}// отправляем данные на клиентскую сторону 
	//вывод списка ссылок
	fmt.Fprintf(w, "<a href = \"https://www.google.ru\"> Тут можно поискать инфу </a> ")
	fmt.Fprintln(w, "<br />") 
 	fmt.Fprintln(w, "<a href = \"https://acegif.com/ru/bobry-na-gifkah/\"> А тут гифки с бобрами </a> <br /> <br />")
	//вывод формы
 	fmt.Fprintln(w, "<form> <label> Name: </label> <input type=\"text\">")
 	fmt.Fprintln(w, "<label> Password: </label> <input type=\"text\"> <br>")
 	fmt.Fprintln(w, "<button type=\"submit\">Send your message</button> </form>")
	resp, err := http.Get("https://google.com") 
   	if err != nil { 
         fmt.Println(err) 
         return
   	} 
   	defer resp.Body.Close() 
   	io.Copy(os.Stdout, resp.Body)
}

func main() {
	http.HandleFunc("/", HomeRouterHandler) // установим роутер
	err := http.ListenAndServe(":4000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}