package main


import (
	"net"
	"fmt"
	"os"
	"bufio"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Print("Input: ")
		msg, _ := reader.ReadString('\n')  //读到换行
		msg = strings.TrimSpace(msg)
		//fmt.Scanln(&msg)
		if msg == "exit"{
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}