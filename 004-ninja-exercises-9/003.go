package main

//Race conditions
import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(100)
	counter := 0

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			fmt.Println("The current index:", i)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count", counter)

}
