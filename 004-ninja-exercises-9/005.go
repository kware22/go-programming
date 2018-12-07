package main

//Atomic
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))

			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end count", counter)

}
