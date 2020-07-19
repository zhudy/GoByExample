package main

import(
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int){
	for j := range(jobs){
		fmt.Println("worker ", id, " started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, " done job", j)
		results<- j * 2
	}
}

func main() {
	const nums = 5
	jobs := make(chan int, nums)
	results := make(chan int, nums)

	for i := 0; i < 3; i++{
		go worker(i, jobs, results)
	}

	for j := 0; j < nums; j++{
		jobs <- j
	}
	close(jobs)

	for k := 0; k < nums; k++{
		<-results
	}


}

