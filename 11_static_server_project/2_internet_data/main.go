// Pull down data from te internet
package main

import (
	"fmt"
	"net/http"
)

// The handle func has been pulled out of main() for ease of readability. Notice
// you can use html when you print line.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>HOME</h1>")
}

// Create handlers for additional pages
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABOUT")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CONTACT")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/contact/", contactHandler)

	fmt.Printf("Listening on port 9000\n")
	http.ListenAndServe(":9000", nil)
}
