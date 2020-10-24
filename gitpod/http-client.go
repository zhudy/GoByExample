package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

// net/http client

func main(){
	resp, err := http.Get("http://127.0.0.1:8080/f2/?name=张三&age=18")
	if err != nil{
		fmt.Printf("get url failed, err: %v\n", err)
		return
	}
	//从resp中将服务器返回的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Printf("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}