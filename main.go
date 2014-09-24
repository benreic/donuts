package main

import (
    "fmt"
    "net/http"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love the home page!")
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
    fmt.Fprintf(w, "Hi there, I love %s!", name)
}

func main() {

	r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/{name}", NameHandler)
    //r.HandleFunc("/products", ProductsHandler)
    //r.HandleFunc("/articles", ArticlesHandler)

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
