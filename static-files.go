package main

import "net/http"

func main() {
    fs := http.FileServer(http.Dir("."))
    http.Handle("/ical/", http.StripPrefix("/ical/", fs))

    http.ListenAndServeTLS(":1234", "icalexample.com+3.pem", "icalexample.com+3-key.pem", nil)
}