package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limit := 1
	burst := 5
	limiter := rate.NewLimiter(rate.Limit(limit), burst)
	//ctx := context.Background()
	i := 0
	for {
		if limiter.Allow() {
			fmt.Println("Do smth...", time.Now())
			//limiter.Wait(ctx)

			fmt.Println(i, " - ", time.Now())
			i++
		} else {
			fmt.Println("Not allowed", " - ", time.Now())
			time.Sleep(time.Second / 3)
		}
	}
}
