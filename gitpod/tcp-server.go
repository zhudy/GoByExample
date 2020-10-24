//https://www.bilibili.com/video/BV16E411H7og?p=113
package main

import (
	"net"
	"fmt"
)
// tcp server
func processConn(conn net.Conn){
		//3. talk with client
		var tmp [128]byte
		defer conn.Close()
		for{
			n, err := conn.Read(tmp[:])
			if err != nil {
				fmt.Println("read from conn failed, err:", err)
				return
			}

			fmt.Println(string(tmp[:n]))
		}
}

func main(){
	//1. local port
	listener, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil{
		fmt.Println("start tcp server on port 20000 failed, err:", err)
		return
	}
	//2. wait for connection
	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("accept failed, err:", err)
			return
		}
		go processConn(conn)
	}
	defer listener.Close()
}

