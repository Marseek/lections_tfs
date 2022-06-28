package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	types2()
}

func types() {
	var a float64

	a = 234.1
	fmt.Println(a)

	fmt.Println(len("Hello World"))
	fmt.Println(len("привет мир"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	for _, val := range "При" {
		fmt.Printf("%T\n", val)
	}
	fmt.Println(hex.Dump([]byte("При")))
}

func types2() {
	const (
		aaa = iota + 3
		bbb
		ccc     = 6493247042348742047206345634563456345623874264455267669785693
		fds int = 234232
	)
	fmt.Printf("%d %d\n\n", aaa, bbb)

	ss := []int{1, 2, 3, 4, 5}
	a := [...]int{1, 2, 3, 4, 5}

	//sss := a[2:3]
	fmt.Printf("%T\n", a)
	ss = append(ss, 1, 3, 4, 5)
	fmt.Println(ss, a)

	b := map[string]int{}
	b["lsdj"] = 234
	b["hello"] = 234
	//delete(b, "hello")
	for k, v := range b {
		fmt.Println(k, v)
	}
	s := make(map[string]int, 10)

	//len(b)
}
