package main

import "fmt"

func insertSort(sli []int) []int {
	sliLen := len(sli)
	for i := 1; i < sliLen; i++ {
		for j := i-1; j > -1; j-- {
			if sli[j+1] < sli[j] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			} else {
				break
			}
		}
	}

	return sli
}

func main() {
	sli := []int{1, 43, 105, 36, 87, 64, 57, 23, 14, 5, 7, 2, 54, 62, 21, 66, 32, 78, 36, 76}
	sli = insertSort(sli)
	fmt.Println("insertSort =", sli)
}