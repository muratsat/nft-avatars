package main

import (
	"fmt"
	"log"
	"net/http"
	handler "nft-avatars/api"
)

func main() {
	http.HandleFunc("/*", handler.Handler)

	fmt.Println("listening at http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("something went wrong when serving")
	}
}
