package main

import "fmt"

type Snack struct {
	data []string
}

func (s *Snack) Push(new string) {
	s.data = append(s.data, new)
}

func (s *Snack) Pop() string {
	var val string
	ln := len(s.data)
	if ln > 0 {
		val = s.data[ln-1]
		s.data = s.data[:ln-1]
	}

	return val
}

func (s Snack) GetAll() []string {
	return s.data
}

func main() {
	var sn = Snack{}

	sn.Push("first")
	sn.Push("second")
	sn.Push("third")

	fmt.Println(sn.Pop())
	fmt.Println(sn.Pop())
	fmt.Println(sn.Pop())

}
