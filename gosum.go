/*
练习 14.5：gosum.go：用这种习惯用法写一个程序，开启一个协程来计算2个整数的合并等待计算结果并打印出来
*/

package main

import (
	"fmt"
)

func sum(x, y int, c chan int) {
	c <- x+y
}

func main() {
	c := make(chan int)
	go sum(12, 13, c)
	fmt.Println(<-c)
}
