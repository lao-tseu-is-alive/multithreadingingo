package main

import (
	"fmt"
	"sync"
	"time"
)

var mlock = sync.Mutex{}
var cond = sync.NewCond(&mlock)

func runChildThread() {
	mlock.Lock()
	fmt.Println("RunChildThread, lock acquired")
	cond.Signal()
	fmt.Println("RunChildThread, Waiting")
	cond.Wait()
	fmt.Println("RunChildThread, Running")
}

func RunMainThread() {
	mlock.Lock()
	fmt.Println("RunMainThread, lock acquired")
	go runChildThread()
	fmt.Println("RunMainThread, Waiting")
	cond.Wait()
	fmt.Println("RunMainThread, Running")
	cond.Signal()
	time.Sleep(10 * time.Second)
}

func main() {
	RunMainThread()
	/*
			will display :
			RunMainThread, lock acquired
			RunMainThread, Waiting
			RunChildThread, lock acquired
			RunChildThread, Waiting
			RunMainThread, Running


		The mutex lock is never unlocked when the signal is called from the RunMainThread,
		so when the RunChildThread is unblocked with the signal,
		it will try to acquire the mlock but will be blocked doing so.
		so the : "RunChildThread, Running" will never display

	*/
}
