package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	log.Println(key)
	fmt.Fprint(w, key)
}
