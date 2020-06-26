package main

import (
	"fmt"
	"os"
	"time"
)

// chan *item ：既能接收也能发送的一个队列
// chan <- *item :只能发送的一个队列（单向通道）
// <-chan *item :只能接收一个队列（单向通道）

func getInput() {
	tmp := [1]byte{}
	os.Stdin.Read(tmp[:])
	fmt.Println(123)

}
func main() {
	getInput()
	time.Sleep(100 * time.Second)

}
