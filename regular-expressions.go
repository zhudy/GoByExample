package main

import (
    "bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println("r.MatchString ", r.MatchString("peach"))
	fmt.Println("r.FindString ", r.FindString("peach punch"))
	fmt.Println("r.FindStringIndex ", r.FindStringIndex("peach punch"))
	fmt.Println("r.FindStringSubmatch ", r.FindStringSubmatch("peach punch"))
	fmt.Println("r.FindStringSubmatchIndex ", r.FindStringSubmatchIndex("peach punch"))
	fmt.Println("r.FindAllString ", r.FindAllString("peach punch pinch", -1))
	fmt.Println("r.FindAllString ", r.FindAllString("peach punch pinch", 2))
	
	fmt.Println("r.Match ", r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")  // panic instead of return an error, safer to use for global variables
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach two punch three pinch", "<fruit>"))

	in := []byte("a peach two puch three pinch")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}