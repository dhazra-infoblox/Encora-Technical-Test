package main

import (
	"log"
	"sync"
)

const goroutines = 2 //Total number of threads to use, excluding the main() thread

func double(v int) {
	log.Printf("Thread %d returned: %d", v, v*2)
	return
}

func main() {
	var ch = make(chan int, 10) // This number 10 can be anything greater than goroutines
	var wg sync.WaitGroup

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for {
				a, ok := <-ch
				if !ok { // if there is nothing to do and the channel has been closed then end the goroutine
					wg.Done()
					return
				}
				double(a) // call double
			}
		}()
	}

	// Now the jobs can be added to the channel, which is used as a queue
	for i := 0; i < 10; i++ {
		ch <- i // append i to the channel
	}
	close(ch) // This tells the goroutines there's nothing else to do
	wg.Wait() // Wait for the goroutines to finish.
}
