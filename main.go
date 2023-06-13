package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server started....")
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Github"))
	})

	srv := &http.Server{
		Addr:    ":6555",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	timer := time.NewTimer(1 * time.Minute)

	<-timer.C

	// Stop the server gracefully
	if err := srv.Shutdown(nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Server gracefully stopped.")
}
