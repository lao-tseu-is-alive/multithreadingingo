package main

import (
	"fmt"
	"sync"
)

func increment(x *int, id int, wg *sync.WaitGroup) {
	fmt.Printf("#Entering -->increment    id:[%d] x: %v \n", id, *x)
	for i := 0; i < 100; i++ {
		*x += 1
	}
	fmt.Printf("#Going out of <--increment id:[%d] x: %v \n", id, *x)
	wg.Done()
}

func count() {
	wg := sync.WaitGroup{}
	x := 0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&x, i, &wg)
	}
	wg.Wait()
	fmt.Printf("%d\n", x)
}

func main() { //nolint:typecheck
	fmt.Println("This code will have a race condition")
	fmt.Println("Just try : go run -race waitgroup.go ")
	count()
}
