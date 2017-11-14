package handlers

import (
	"fmt"
	//"html/template"
	"io/ioutil"
	"net/http"
)

//Root TODO: should read cockies and redirect to signin/register or blog page
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
