package main

import (
	"fmt"
)

func BubbleSort(byteSli []byte) []byte {
	compareLen := len(byteSli)
	for compareLen > 1 {
		for i := 0; i < compareLen-1; i++ {
			if byteSli[i] > byteSli[i+1] {
				byteSli[i+1], byteSli[i] = byteSli[i], byteSli[i+1]
			}
		}
		compareLen--
	}
	return byteSli
}

func main() {
	byteSli := []byte{14, 61, 12, 59, 3, 15, 48, 5, 23, 36}
	byteSli = BubbleSort(byteSli)
	fmt.Println("BubbleSort =", byteSli)
}
