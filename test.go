package main

import (
	"fmt"
	"os"
	"time"
)

func getInput() {
	tmp := [1]byte{}
	os.Stdin.Read(tmp[:])
	fmt.Println(123)

}
func main() {
	getInput()
	time.Sleep(100 * time.Second)

}
