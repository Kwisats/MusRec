package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	//"html/template"
	"io/ioutil"
	"net/http"
)

const registeredUsers = "users"

//User ...
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login/" {
		fmt.Fprintf(w, "<script>alert(\"page %s not found\")</script>", r.URL.Path)
		return
	}
	if r.Method == http.MethodPost {
		inputEmail := r.FormValue("inputEmail")
		inputPassword := r.FormValue("inputPassword")
		fmt.Println("Login is:", inputEmail)
		fmt.Println("Password is:", inputPassword)
		//TODO: Validation

		file, err := os.Open(registeredUsers)
		defer file.Close()
		if err != nil {
			fmt.Println("file open error:", err)
			return
		}
		reader := bufio.NewReader(file)
		user := User{}
		for {
			line, err := reader.ReadBytes(byte('\n'))
			if err != nil {
				break
			}
			err = json.Unmarshal(line, &user)
			if err != nil {
				fmt.Println("json unmarshal error:", err)
				return
			}
			if user.Login == inputEmail {
				if user.Password == inputPassword {
					//TODO something
					fmt.Println("I know him")
					return
				}
			} else {
				continue
			}
		}
		fmt.Fprintln(w, "<script>alert(\"wrong login or password\")</script>")
	}
	fileContent, err := ioutil.ReadFile("static/pages/signin.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(fileContent)
}
