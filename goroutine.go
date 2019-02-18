package main

import (
	"context"
	"fmt"
	"runtime"
)

func say(s string, cancelFunc context.CancelFunc) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	//runtime.NumCPU()
	//sync.WaitGroup{}
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	c := make(chan int)
	_, confunc := context.WithCancel(context.TODO())
	go say("world", confunc) //开一个新的Goroutines执行
	//say("hello")    //当前Goroutines执行
	<-c
	//time.Sleep(time.Second * 1
}
