package main

import (
	"log"
	"mutants/app/cmd"
	"net/http"
	"os"
)

func main() {
	e := cmd.Run()

	http.Handle("/", e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
