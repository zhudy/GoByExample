package main

import (
	"fmt"
	"os"
)

func main() {
	//f := createFile("/tmp/defer.txt")
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f, "hello")
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)

	if err != nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File, s string){
	fmt.Println("writing")
	fmt.Fprintf(f, s)
}

func closeFile(f *os.File){
	fmt.Println("closing")
	err := f.Close()

	if err != nil{
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}