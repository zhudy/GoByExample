package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request){
	str, err := ioutil.ReadFile("./hello.txt", )
	if err != nil{
		w.Write([]byte("找不到文件 hello.txt, 无法显示其内容" + fmt.Sprintf("%v", err)))
	}
	//str := `<h1 style="color:red">Hello WWW!</h1>`
	w.Write([]byte(str))
}

func main(){
	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.ListenAndServe(":8080", nil)
}