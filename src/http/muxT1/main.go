package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "home handler")
}

func ProductsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Products handler")
}

func t1() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	http.Handle("/", r)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func t2() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/articles/{category}", ArticlesCategoryHandler)
	http.Handle("/", r)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func main() {
	t2()
}
