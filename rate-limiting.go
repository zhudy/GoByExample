package main

import(
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++{
		requests <- i
	}
	close(requests)
	limit := time.Tick(200 * time.Millisecond)
	for req := range requests{
		<-limit
		fmt.Println("req", req, time.Now())
	}

	fmt.Println("\n\n")
	burstyLimiter := make(chan time.Time, 3)
	for j := 0; j < 3; j++{
		burstyLimiter <- time.Now()
	}

	go func(){
		for t := range time.Tick(1000 * time.Millisecond){
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for j := 0; j < 5; j++{
		burstyRequest <- j
	}
	close(burstyRequest)

	for req := range burstyRequest {
		<- burstyLimiter
		fmt.Println("req", req, time.Now())
	}
}