package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("serving on 9999")
	http.Handle("/", http.FileServer(http.Dir("./pages/")))
	http.Handle("/page2", http.FileServer(http.Dir("./pages/")))
	http.ListenAndServe(":9999", nil)
}
