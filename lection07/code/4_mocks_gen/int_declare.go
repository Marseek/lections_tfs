package main

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

// use ./mockery --inpackage --name=Connection to generate mock connection
type Connection interface {
	GetConn() int
	SetConn(int) bool
}

func main() {
	mockConn := new(MockConnection)
	mockConn.Mock.On("GetConn").Return(0)
	i := mockConn.GetConn()
	mockConn.GetConn()
	fmt.Println(i)

	mockConn.Mock.On("SetConn", mock.Anything).Return(true)
	b := mockConn.SetConn(i)
	fmt.Println(b)
}
