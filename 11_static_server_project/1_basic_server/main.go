// An initial implenmentation of a server to demonstrate the concepts

package main

// import the standard go package net/http to handle http requests and
// responses.
import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new router.
	// args:
	//		The path you wish to be written on.
	//		func, response writer, request listner
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// What you want the response writer to do.
		// In this case write "Hello Go" to the page. You could also ask it to
		// serve up html, xml, css etc.
		w.Write([]byte("Hello Go"))

	})

	// Add a message in the terminal so we can see it's running. At this point
	// you will receive a 404 as it points nowhere.
	fmt.Println("server running on port 8080")

	// Launch go's default http server
	// args:
	//		given address (on our selected port, localhost:8080 in this case).
	//		handler (usually `nil` to use DefaultServeMux. Handle and HandleFunc
	//		add handlers to DefaultServeMux).
	http.ListenAndServe(":8080", nil)
}
