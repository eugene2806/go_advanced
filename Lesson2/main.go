package main

import (
	"fmt"
	"log"
	"net/http"
)

//w - responseWriter (куда писать ответ)
//r - request (откуда брать запрос)
// Обработчик

func GetGreat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, I'm New Web Server")
}

// Товарищ, который выбирает нужный обработчик в зависимости от адреса, по которому пришел запрос

func RequestHandler() {
	http.HandleFunc("/", GetGreat)               // Если придет запрос по адресу / то вызывает GetGreat()
	log.Fatal(http.ListenAndServe(":8080", nil)) // Запускаем веб сервер в режиме слушания
}
func main() {
	RequestHandler()
}
