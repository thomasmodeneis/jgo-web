package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	p := os.Getenv("PORT")
	if p == "" {
		p = "3006"
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening -> "+p)
	http.ListenAndServe(":"+p, nil)
}
