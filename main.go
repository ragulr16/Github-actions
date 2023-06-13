package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server started....")
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Github"))
	})

	if err := http.ListenAndServe(":6565", r); err != nil {
		log.Fatal(err)
	}
}
