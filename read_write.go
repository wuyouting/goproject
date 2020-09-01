package main

import (
	"fmt"
	"time"
)

//读比写多的时候使用读写锁
var x int64

func read() {
	fmt.Println(x)
	time.Sleep(time.Microsecond * 10)
}
func write() {
	x = x + 1
	time.Sleep(time.Microsecond * 50)

}
func main() {
	now := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
}
