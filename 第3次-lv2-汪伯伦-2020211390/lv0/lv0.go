//// @Title : lv0
//// @Description :改写加锁
//// @Author : MX
//// @Update : 2021/11/19 20:24
//
package main

import (
	"fmt"
)

var (
	myres = make(map[int]int, 20)
	//mu sync.Mutex
	ch = make(chan struct{},1)
)
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//mu.Lock()
	<- ch
	myres[n] = res
	//mu.Unlock()
	ch <- struct{}{}
}

func main() {
	ch <- struct{}{}
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}

	//mu.Lock()
	<- ch
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
	//mu.Unlock()
	ch <- struct{}{}
}

