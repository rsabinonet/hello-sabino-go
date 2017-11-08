package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	} else {
		fmt.Println("Running in PORT:", port)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello do Sabino que ama a Pilouita !")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
