package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/imadedwisuryadinta/dts-microservice/auth-service/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/handle-admin", http.HandlerFunc(handler.ValidateAuth))

	fmt.Println("Auth service listen on port 8001")
	log.Panic(http.ListenAndServe(":8001", router))
}
