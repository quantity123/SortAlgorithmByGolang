package main

import "fmt"

func SlectSort(byteSlice []byte) {
	byteSliceLen := len(byteSlice)
	for i:=0; i<byteSliceLen-1; i++ {
		for j:=i+1; j<byteSliceLen; j++ {	//每次j+1让开头数与后一位开始比较
			if byteSlice[i] > byteSlice[j] {
				byteSlice[j], byteSlice[i] = byteSlice[i], byteSlice[j]	//这个传值非常强，不用写转换函数
			}
		}
	}
}

func main() {
	byteSlice := []byte{1, 43, 54, 62, 21, 66, 32, 78, 36, 59, 25, 36, 72, 14, 23, 11, 14, 6, 9, 76}
	SlectSort(byteSlice)
	fmt.Println(byteSlice)
}
