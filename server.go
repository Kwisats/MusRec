package main

import (
	"fmt"
	"net/http"
	"time"
	h "training/handlers"
)

func runServer(addr string) {
	mux := http.NewServeMux()

	//import static data
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
	mux.Handle("/styles/", stylesHandler)
	mux.Handle("/js/", jsHandler)
	mux.Handle("/img/", imgHandler)

	//create pages
	mux.HandleFunc("/", h.Root)
	mux.HandleFunc("/test/", h.Test)
	mux.HandleFunc("/login/", h.Login)

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
	fmt.Println("Main starts")
	runServer(":8080")
}
