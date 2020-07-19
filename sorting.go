package main

import(
	"fmt"
	"sort"
)

func main() {

	s := []string{"c","b","a"}
	sort.Strings(s)
	fmt.Println(s)

	i := []int{3,7,5}
	fmt.Println(i)
	sorted := sort.IntsAreSorted(i)
	fmt.Println("sorted: ", sorted)

	sort.Ints(i)
	fmt.Println(i)
	sorted = sort.IntsAreSorted(i)
	fmt.Println("sorted: ", sorted)
}