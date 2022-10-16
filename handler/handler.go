package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(
		path.Join("views", "index.html"),
		path.Join("views", "layout.html"),
	)
	if err != nil {
		log.Println(err)
		// log.Fatalln(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"title":   "I'm Learning Golang Web",
		"content": "I'm learning Golang Web with Agung Setiawan",
	}

	// w.Write([]byte("Welcome to Home"))
	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halo dunia, saya sedang belajar Golang web"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from Nintendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Product Page"))
	// fmt.Fprintf(w, "Product Page: %d", idNum)

	tmpl, err := template.ParseFiles(
		path.Join("views", "product.html"),
		path.Join("views", "layout.html"),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"title":   "I'm Learning Golang Web | Product",
		"product": entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3},
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}
