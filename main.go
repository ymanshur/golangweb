package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	}

	mux.HandleFunc("/", homeHandler) // root, default route
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) { // anonymous function
		w.Write([]byte("Profile Page"))
	})

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
