package main

import "fmt"

func BuffSender() {
	channel := make(chan string, 3)
	fmt.Println("Sending ONE")
	channel <- "ONE"
	fmt.Println("Sending TWO")
	channel <- "TWO"
	fmt.Println("Sending THREE")
	channel <- "THREE"
	fmt.Println("Done")
}
func runConsumer(channel chan string) {
	msg := <-channel
	fmt.Println("Consumer, received", msg)
	channel <- "Bye"
}

func RunProducer() {
	channel := make(chan string)
	go runConsumer(channel)
	fmt.Println("Producer Sending Hello")
	channel <- "Hello from Producer"
	fmt.Println("Producer, received", <-channel)
}

func main() {
	RunProducer()

	BuffSender()
}
