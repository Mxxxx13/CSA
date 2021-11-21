// @Title : lv2
// @Description :并发素数筛
// @Author : MX
// @Update : 2021/11/20 16:13 

package main

import "fmt"

// generateNatural 返回生成自然数序列的管道
func generateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// primeFilter 素数筛,每次删除能被素数(prime)整除的数
func primeFilter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := generateNatural()
	for i := 0; ; i++ {
		prime := <-ch
		if prime > 50000 {
			break
		}
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = primeFilter(ch, prime)
	}
}
