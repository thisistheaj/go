package main

import "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!") 
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Serving on port http://localhost:8000")
    http.ListenAndServe(":8000", nil)
}
