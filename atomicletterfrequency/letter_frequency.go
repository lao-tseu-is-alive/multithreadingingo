package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

var lock = sync.Mutex{}

func countLetters(url string, frequency *[26]int32, wg *sync.WaitGroup) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	for i := 0; i <= 20; i++ {
		for _, b := range body {
			c := strings.ToLower(string(b))
			//lock.Lock()
			index := strings.Index(allLetters, c)
			if index >= 0 {
				//*frequency[c] += 1  // ~25sec
				// by using atomic variables we reduce time execution to ~ 3sec (instead of 25)
				// atomic variables are a quick win for this kind of problems
				atomic.AddInt32(&frequency[index], 1)
			}
			//lock.Unlock()
		}
	}
	wg.Done()
}

func main() {
	var frequency [26]int32
	wg := sync.WaitGroup{}

	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s\n", elapsed)
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}
