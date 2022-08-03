package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		fmt.Println("Blue: Acquiring lock1")
		lock1.Lock()
		fmt.Println("Blue: Acquiring lock2")
		lock2.Lock()
		fmt.Println("Blue: Both locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: Locks Released")
	}
}

func redRobot() {
	for {
		fmt.Println("Red: Acquiring lock2")
		lock2.Lock()
		fmt.Println("Red: Acquiring lock1")
		lock1.Lock()
		fmt.Println("Red: Both locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Red: Locks Released")
	}
}

// https://yourbasic.org/golang/detect-deadlock/
// CONCLUSION DO NOT USE SAME LOCK ON MANY GOROUTINES WITHOUT ANOTHER WAY TO COORDINATE ACCESS  BETWEEN GO ROUTINES !!!
// https://go.dev/doc/articles/race_detector does not detect deadlock
// maybe you can try using  https://github.com/sasha-s/go-deadlock
// introduce a resource hierarchy : check deadlocks_train example
func main() {
	go redRobot()
	go blueRobot()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
