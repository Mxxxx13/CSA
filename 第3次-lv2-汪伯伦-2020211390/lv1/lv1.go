// @Title : lv1
// @Description :channel控制输出打印
// @Author : MX
// @Update : 2021/11/19 20:24 

package main

import (
	"fmt"
	"sync"
)

func printA(ch1,ch2 chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		<-ch1
		fmt.Printf("%d: A ",i + 1)
		wg.Done()
		ch2 <- struct{}{}

	}
}

func printB(ch2,ch3 chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		<-ch2
		fmt.Printf("B ")
		wg.Done()
		ch3 <- struct{}{}

	}
}

func printC(ch3,ch1 chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		<-ch3
		fmt.Printf("C\n")
		wg.Done()
		ch1 <- struct{}{}
	}
}

func main() {
	
	var (
		ch1 = make(chan struct{}, 1)
		ch2 = make(chan struct{}, 1)
		ch3 = make(chan struct{}, 1)
		wg  sync.WaitGroup
	)
	defer wg.Wait()
	ch1 <- struct{}{}
	go printA(ch1,ch2,&wg)
	go printB(ch2,ch3,&wg)
	go printC(ch3,ch1,&wg)
}
