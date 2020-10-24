package main

import (
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request){
	str := `<h1 style="color:red">Hello WWW!</h1>`
	w.Write([]byte(str))
}

func main(){
	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.ListenAndServe(":8080", nil)
}