package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	UserService "main/modules/user/services"
	"main/repository/mysql"
	"net/http"
)

var mysqlRepo *mysql.MySqlDB

func init() {
	mysqlRepo = mysql.New()
}

func main() {
	http.HandleFunc("/api/health", healthCheck)
	http.HandleFunc("/api/register-by-phone", registerByPhone)

	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerByPhone(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(writer, `{"error" : "Method not allowed"}`)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}
	defer request.Body.Close()

	var ureq UserService.RegisterRequest
	err = json.Unmarshal(body, &ureq)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	userService := UserService.New(mysqlRepo)

	_, err = userService.Register(ureq)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, `{"message" : "User registered successfully"}`)
}

func healthCheck(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, `{"message":"Every thing is working bro"}`)
}
