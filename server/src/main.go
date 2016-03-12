package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

}

func search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form["key"]

}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/search", search)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
