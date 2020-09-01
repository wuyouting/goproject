package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

//定义一个函数 对全局的变量X做
func add() {

	for i := 0; i < 5000; i++ {
		lock.Lock() //把厕所门锁上
		//1,从内存中找到x的值
		//2，执行+1 操作
		//3，把结果复制给x写到内存中
		x = x + 1
		lock.Unlock() //把厕所门打开
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
