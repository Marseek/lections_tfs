package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	//chanEx()

	//defer log.Println("Defer 1")
	//defer log.Println("Defer 2")
	//defer log.Println("Defer 3")

	wg_wait_example()
}

func chanEx() {
	ch := make(chan int, 1)

	ch <- 0

	go readFromCh(ch)
	log.Println("write to channel")

	time.Sleep(time.Second)
}

func readFromCh(ch <-chan int) {
	//time.Sleep(time.Second * 2)
	if val, ok := <-ch; ok {
		log.Println("value from channel: ", val)
		return
	}
}

func wg_wait_example() {
	var wc sync.WaitGroup
	wc.Add(5)
	m := make(chan string, 5)

	for i := 0; i < 5; i++ {

		go func(mm chan<- string, i int, group *sync.WaitGroup) {

			defer wc.Done()
			mm <- fmt.Sprintf("Goroutine %d", i)

		}(m, i, &wc)

	}
	wc.Wait()
	close(m)
	for q := range m {
		fmt.Println(q)
	}

}
