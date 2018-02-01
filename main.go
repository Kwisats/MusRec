package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"log"
	h "MusRec/pages"
)

func runServer(addr string) {
	mux := http.NewServeMux()

	//import static data and templates
	stylesHandler := http.StripPrefix(
		"/styles/",
		http.FileServer(http.Dir("./static/styles")),
	)
	jsHandler := http.StripPrefix(
		"/js/",
		http.FileServer(http.Dir("./static/js")),
	)
	imgHandler := http.StripPrefix(
		"/img/",
		http.FileServer(http.Dir("./static/img")),
	)
	templatesHandler := http.StripPrefix(
		"/img/",
		http.FileServer(http.Dir("./static/img")),
	)
	mux.Handle("/styles/", stylesHandler)
	mux.Handle("/js/", jsHandler)
	mux.Handle("/img/", imgHandler)
	mux.Handle("/templates/", templatesHandler)

	//create pages
	mux.HandleFunc("/", h.Root)

	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	fmt.Println("Starting server at", addr)
	server.ListenAndServe()
}

func main() {
	port := os.Getenv("PORT")
	port = "8080"
	if port == "" {
        log.Fatal("$PORT must be set")
    }
	runServer(":" + port)
}
