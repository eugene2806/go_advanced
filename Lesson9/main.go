package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

var dataNap map[int]int

func init() {
	dataNap = make(map[int]int)
}
func main() {
	http.HandleFunc("/factorial", HandlerFactorial)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	if result, ok := dataNap[n]; ok {
		return result
	}
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	dataNap[n] = ans
	return ans
}

func HandlerFactorial(writer http.ResponseWriter, rec *http.Request) {
	// http://localhost:8080/factorial?num=6
	num := rec.FormValue("num")
	n, err := strconv.Atoi(num)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	io.WriteString(writer, strconv.Itoa(factorial(n)))
}
