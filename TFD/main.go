package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/factorial", HandleFactorial)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func factorial(num int) int {
	var result int = 1
	for i := 1; i <= num; i++ {
		result = result * i
	}
	return result
}

func HandleFactorial(writer http.ResponseWriter, request *http.Request) {
	num, err := strconv.Atoi(request.FormValue("num"))
	if err != nil {
		http.Error(writer, err.Error(), 404)
		return
	}
	io.WriteString(writer, strconv.Itoa(factorial(num)))
}
