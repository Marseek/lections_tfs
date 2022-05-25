package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Second)
		ch := ctx.Done()
		for {
			select {
			case answ := <-ch:
				fmt.Printf("%T\n", answ)
				return
			case <-ticker.C:
				fmt.Println(ctx.Err())
			}
		}
	}()

	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println(ctx.Err())
	fmt.Println(ctx.Err())
	wg.Wait()
}
