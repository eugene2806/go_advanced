package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	// Порт запуска приложения
	port string = "8080"
	// База данных
	bd []Pizza
)

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 25,
		Price:    500,
		Title:    "Pepperoni",
	}

	pizza2 := Pizza{
		ID:       2,
		Diameter: 22,
		Price:    550,
		Title:    "BBQ",
	}

	pizza3 := Pizza{
		ID:       3,
		Diameter: 20,
		Price:    300,
		Title:    "Margaritta",
	}

	bd = append(bd, pizza1, pizza2, pizza3)
	log.Println("Pizza bd loaded")
}

// Наша модель

type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

func FindPizzaByID(id int) (Pizza, bool) {
	var pizza Pizza
	var find bool
	for _, p := range bd {
		if p.ID == id {
			pizza = p
			find = true
			break
		}
	}

	return pizza, find
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	//Прописывать хедеры
	writer.Header().Set("Content-Type", "application/json")

	log.Println("Get info all pizzas in database")
	writer.WriteHeader(200)

	json.NewEncoder(writer).Encode(bd) //Сериализация + Запись writer

}

func GetPizzaByID(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// Считаем ID из строки запроса и конвертируем его в int
	vars := mux.Vars(request) // Получаем мапу типа {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Error converting id to int => Error:", err)
		msg := ErrorMessage{Message: "Do not use ID not supported int casting"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	log.Println("Trying to send to client pizza with ID: ", id)
	pizza, ok := FindPizzaByID(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		msg := ErrorMessage{Message: "Pizza not found"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
	}
}

func main() {
	log.Println("Trying to start REST API Pizza")

	// Инициализируем маршрутизатор
	router := mux.NewRouter()

	// Если на вход пришел запрос /pizzas
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")

	// Если на вход пришел запрос /pizza/{id}
	router.HandleFunc("/pizza/{id}", GetPizzaByID).Methods("GET")

	log.Println("Router configured successfully")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
