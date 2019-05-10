/*
练习 14.1：channel_block3.go：写一个通道证明它的阻塞性，开启一个协程接收通道的数据，
持续 15 秒，然后给通道放入一个值。在不同的阶段打印消息并观察输出
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(){
		x := <-c
		fmt.Println("received", x)
	}()
	go func(){
		time.Sleep(time.Second*15)
		fmt.Println("sending", 10)
		c <- 10
		fmt.Println("sent", 10)
	}()
	time.Sleep(time.Second*20)
}
