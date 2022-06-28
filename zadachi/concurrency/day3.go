package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	fmt.Println("2 init!")
}
func init() {
	fmt.Println("first init!")
}

func main() {
	atomics()

}

func atomics() {
	var a int64 = 10
	var b atomic.Value

	//b.Store("sldkfj")
	//b.Store("second string")
	//b.Store("third string")
	b.Store(4)

	fmt.Println(b.Load())
	atomic.AddInt64(&a, 32)
	atomic.AddInt64(&a, 32)
	fmt.Println(a)

	wg := sync.WaitGroup{}
	wg.Add(5)

	smp := sync.Map{}
	smp.Store("sldjf", "lsdkfj")
	fmt.Println(smp.Load("slddjf"))
	fmt.Println(smp.Load("sldjf"))

	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	type message struct {
		cond *sync.Cond
		msg  string
	}

	msg := message{
		cond: sync.NewCond(&sync.Mutex{}),
	}
	cnd := msg.cond
	cnd.Wait()
	cnd.L.Lock()
	// Тестовая платформа i5 3,8 ГГц 4 ядра

	pool := sync.Pool{New: func() interface{} {
		return make([]string, 3, 6)
	}}
	//s := "Hello World"
	//s2 := "Hello World2222"
	s3 := []string{"lsdk", "lsdkld"}
	//pool.Put(s)
	//pool.Put(s2)
	pool.Put(s3)
	s4 := pool.Get().([]string)
	s5 := pool.Get().([]string)
	fmt.Println(len(s4), cap(s4), s4)
	fmt.Println(len(s5), cap(s5), s5)
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())

	var sMap sync.Map
	sMap.Store("first", "first_value")
	sMap.Store("second", "second_value")
	val, _ := sMap.Load("second")
	fmt.Println(val)

}

func main3() {
	c := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			c <- i
		}
		close(c)
	}()
	for {
		i, ok := <-c
		if !ok {
			fmt.Println("chanel closed")
			break
		}
		fmt.Println(i)

	}
}

func main2() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
