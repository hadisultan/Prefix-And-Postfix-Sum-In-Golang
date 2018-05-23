package main

import (
    "fmt"
)

func PostfixSum(data []int, parent chan int) {
	if len(data) > 1 {
		mid := len(data)/2
		left := make(chan int)
		right := make(chan int)
		go PostfixSum(data[:mid], left)
		go PostfixSum(data[mid:], right)
		rightSum := <-right
		parent<- <-left + rightSum
		fromRight := <-parent
        right<- fromRight
		left<- fromRight + rightSum
		parent<- <-left + <-right
	} else {
		parent<- data[0]
		data[0] += <-parent
		parent<- 0
	}
}

func main() {
	data:= []int{2, 6, 1, 7, 3, 0, 12, 8}
	ch := make(chan int)
	fmt.Println(data)
	go PostfixSum(data, ch)
	<-ch
	ch<-0
	<-ch
	fmt.Println(data)
}






