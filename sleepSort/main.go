package main

import (
	"fmt"
	"time"
)

func main() {
	tab := []int{23, 15, 46, 21, 45, 83, 33, 1, 3, 0, 5}

	ch := make(chan int)
	for _, value := range tab {
		go func(val int) {
			time.Sleep(time.Duration(val) * time.Millisecond)//10000000 //数值来sleep睡眠线程延长返回时间
			fmt.Println("var =", val)
			ch <- val
		}(value)
	}

	for i, _ := range tab {
		tab[i] = <-ch
	}

	fmt.Println("tab chan =", tab)
}
