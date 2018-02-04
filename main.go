package main

import (
	t "MusRec/fileHandler"
	h "MusRec/handlers"
	"fmt"
	"log"
	"net/http"
	"time"
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
	//port := os.Getenv("PORT")
	port := "8080"
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fileName := "ostan.mp3"
	err := t.ToRaw(fileName)
	check(err)
	runServer(":" + port)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
