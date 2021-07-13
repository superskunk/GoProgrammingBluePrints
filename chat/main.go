package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com-superskunk/superskunk/GoProgrammingBluePrints/trace"
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
	t.templ.Execute(w, r)
}

func getTracer(tracing bool) trace.Tracer {
	if tracing {
		return trace.New(os.Stdout)
	}
	return trace.Off()
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	var tracing = flag.Bool("t", false, "Activates logs")
	flag.Parse()
	r := newRoom()
	r.tracer = getTracer(*tracing)
	chatTemplate := &templateHandler{filename: "chat.html"}
	http.Handle("/", chatTemplate)
	http.Handle("/room", r)
	go r.run()
	// Start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
