package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Println("Listening on 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Filed to run server:", err)
	}
}
