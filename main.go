package main

import (
	"log"
	"net/http"
	"os"
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

	// Use a boolean flag to track whether the application should exit
	// Create a channel to receive an exit signal
	exit := make(chan struct{})

	// Set a timer for 2 minutes
	timer := time.NewTimer(2 * time.Minute)

	// Start a goroutine to wait for the timer and send an exit signal
	go func() {
		<-timer.C
		close(exit)
	}()

	// Wait for the exit signal or the application to be forcefully exited
	select {
	case <-exit:
		// Stop the server gracefully
		if err := srv.Shutdown(nil); err != nil {
			log.Fatal(err)
		}
		log.Println("Server gracefully stopped.")
	case <-time.After(0):
	}

	// Exit the application without displaying the warning message
	os.Exit(0)
}
