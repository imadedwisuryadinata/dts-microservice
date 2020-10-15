package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/imadedwisuryadinta/dts-microservice/menu-service/handler"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("menu service listen on port : 8000")
	log.Panic(http.ListenAndServ("8000", router))
}
