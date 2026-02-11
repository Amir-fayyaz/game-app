package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/health", healthCheck)
	http.HandleFunc("/api/register-by-phone", registerByPhone)

	http.ListenAndServe(":8080", nil)
}

func registerByPhone(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	if request.Method != http.MethodPost {
		fmt.Println(request.Method)
		fmt.Fprint(writer, "Method not allowed")
	}

	fmt.Println("Here !")
}

func healthCheck(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Every thing is working bro")
}
