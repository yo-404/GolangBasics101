package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myChannel := make(chan int, 1)
	wg := &sync.WaitGroup{}
	// myChannel <- 4
	// fmt.Println(<-myChannel)
	wg.Add(2)
	// creating recieve only using <-chan
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, IschannelOpen := <-myChannel
		fmt.Println(IschannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myChannel)
		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)
	// creating send only using <-chan
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 0
		myChannel <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)
	wg.Wait()
}
