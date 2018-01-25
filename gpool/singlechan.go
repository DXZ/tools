package gpool

import (
	"fmt"
)

const (
	Max = 10
)

type SinglePool struct {
	queue chan int
}

func NewSinglePool(quNum int) *SinglePool {
	return &SinglePool{
		queue: make(chan int, quNum),
	}
}

func StartWork() {

}

func main() {
	fmt.Println("start...")
}
