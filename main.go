package main

import (
	"fmt"
	"log"
	"net/http"
	handler "nft-avatars/api"
)

func main() {
	http.HandleFunc("/", handler.Handler)

	fmt.Println("listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("something went wrong when serving")
	}
}
