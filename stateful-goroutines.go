package main
import(
	"fmt"
	"sync/atomic"
	"time"
	"math/rand"
)

type readOp struct{
	key int
	resp chan int
}
type writeOp struct{
	key int
	value int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func(){
		var state = make(map[int]int)
		for{
			select{
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++{
		go func(){
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++{
		go func(){
			for {
				write := writeOp{
					key : rand.Intn(5),
					value : rand.Intn(100),
					resp : make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println(readOpsFinal)
	fmt.Println(writeOpsFinal)
}