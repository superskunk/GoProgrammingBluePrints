package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// The first time ServeHTTP is called, it load the template ("templates/t.filename") and compile it
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	// Executes template and write the output to the specified http.ResponseWirter method
	t.templ.Execute(w, nil)
}

func main() {

	chatTemplate := &templateHandler{filename: "chat.html"}

	http.Handle("/", chatTemplate)

	// Start the web server

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
