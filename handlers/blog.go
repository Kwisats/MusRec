package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Blog ..
func Blog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog/" {
		fmt.Fprintf(w, "<script>alert(\"page %s not found\")</script>", r.URL.Path)
		return
	}
	fileContent, err := ioutil.ReadFile("static/pages/blog.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(fileContent)
}
