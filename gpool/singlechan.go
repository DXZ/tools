package gpool

import (
	"fmt"
	oldpool "ginfoxer/tools/gpool"
	"time"
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

// func Test(i int) {
// 	fmt.Println("--------doing job before---------", i)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("--------doing job after---------", i)
// }

// func main() {
// 	fmt.Println("start...")
// 	num := 10
// 	fgpool := oldpool.New(num)
// 	{
// 		for i := 0; i < num+10; i++ {
// 			fgpool.Add(1)
// 			go func(i int) {
// 				Test(i)
// 				fgpool.Done()
// 			}(i)
// 		}
// 	}
// 	fgpool.Wait()
// }
