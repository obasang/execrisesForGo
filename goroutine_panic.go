/*
练习 14.7：

a）在练习 5.4 的 for_loop.go 中，有一个常见的 for 循环打印数字。
在函数 tel 中实现一个 for 循环，用协程开始这个函数并在其中给通道发送数字。
main() 线程从通道中获取并打印。不要使用 time.Sleep() 来同步：goroutine_panic.go
b）也许你的方案有效，可能会引发运行时的 panic：throw:all goroutines are asleep-deadlock! 
为什么会这样？你如何解决这个问题？goroutine_close.go
c）解决 a）的另外一种方式：使用一个额外的通道传递给协程，然后在结束的时候随便放点什么进去。
main() 线程检查是否有数据发送给了这个通道，如果有就停止：goroutine_select.go
*/

package main

import (
	"fmt"
)

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
}

func main() {
	var ok = true
	ch := make(chan int)

	go tel(ch)
	for ok {
		i := <-ch
		fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	}
}
