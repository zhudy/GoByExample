package main

import(
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int{
	return len(s)
}
func (s byLength) Swap(i, j int){
	// tmp := s[i]
	// s[i] = s[j]
	// s[j] = tmp
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool{
	return len(s[i]) < len(s[j])
}

func main() {

	fruits := []string{"apple", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}