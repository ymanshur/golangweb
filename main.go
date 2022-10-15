package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler) // root, default route
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)

	log.Println("Stating web on port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to Home"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halo dunia, saya sedang belajar Golang web"))
}

func marioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from Nintendo"))
}
