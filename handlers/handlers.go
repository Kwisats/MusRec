package handlers

import (
	"fmt"
	//"html/template"
	"io/ioutil"
	"net/http"
)

//Handler should be used like this
/*
	rootHandler := &Handler{Name: "root"}
	mux.Handle("/", rootHandler)
*/
type Handler struct {
	Name string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Some Actions
}

//Test ...
func Test(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test/" {
		fmt.Fprintf(w, "<script>alert(\"page %s not found\")</script>", r.URL.Path)
		return
	}
	fileContent, err := ioutil.ReadFile("static/pages/test.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(fileContent)
}

//Root ...
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintf(w, "<script>alert(\"page %s not found\")</script>", r.URL.Path)
		return
	}
	fileContent, err := ioutil.ReadFile("static/pages/root.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(fileContent)
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login/" {
		fmt.Fprintf(w, "<script>alert(\"page %s not found\")</script>", r.URL.Path)
		return
	}
	fileContent, err := ioutil.ReadFile("static/pages/login.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(fileContent)
}
