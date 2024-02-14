package main

import (
	"fmt"
	"log"
	"net/http"
	handler "nft-avatars/api"
)

func main() {
	fmt.Println("Hello world")
	http.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("something went wrong when serving")
	}
}
