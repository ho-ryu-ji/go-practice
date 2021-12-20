package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, World!")
			log.Printf("%+v¥n", r)
		})
	http.ListenAndServe(":8080", nil)
}
