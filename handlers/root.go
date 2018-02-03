package handlers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//Root is handle function of route page
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Fatal("unknown way: " + r.URL.Path)
		return
	}
	if r.Method == "POST" {
		src, hdr, err := r.FormFile("my_file")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer src.Close()

		dst, err := os.Create(filepath.Join("musicFiles", hdr.Filename))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer dst.Close()
		io.Copy(dst, src)

		//this is the example of using bash-script in go
		//out, err := exec.Command("bashScript").Output()
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("output is %s\n", out)
	}
	uploadForm, err := ioutil.ReadFile("templates/upload.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(uploadForm)
}
