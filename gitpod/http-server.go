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

func f2(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL)
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
	//w.flush()
}

func main(){
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/f2/", f2)
	http.ListenAndServe(":8080", nil)
}