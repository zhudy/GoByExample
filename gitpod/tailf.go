package main

import (
	"github.com/hpcloud/tail"
	"fmt"
)

func main(){
	t, err := tail.TailFile("./hello.txt", tail.Config{Follow: true})
	if err != nil{
		fmt.Println("TailFile failed err: %v", err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
