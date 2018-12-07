package main

//goRoutines
import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("This is foo!")
	wg.Done()
}

func bar() {
	fmt.Println("This is barrrr!")
	wg.Done()
}

func main() {

	wg.Add(2)
	fmt.Println("1. Current amount of go routines: ", runtime.NumGoroutine())
	go foo()
	fmt.Println("2. Current amount of go routines: ", runtime.NumGoroutine())
	go bar()
	fmt.Println("3. Current amount of go routines: ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("about to exit")

}
