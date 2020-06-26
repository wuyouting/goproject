package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

//定义一个函数 对全局的变量X做
func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
