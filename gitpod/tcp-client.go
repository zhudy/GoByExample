package main


import (
	"net"
	"fmt"
	"os"
)

//tcp client
func main(){
	//1. connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial 127.0.0.1:20000 failed, err:", err)
		return
	}

	//2. send data
	var msg string
	if len(os.Args) < 2{
		msg = "hello world!"
	}else{
		msg = os.Args[1]
	}
	conn.Write([]byte(msg))
	conn.Close()
}