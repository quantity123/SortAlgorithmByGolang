package main

import "fmt"

func QuickSort(list []byte) (returnList []byte) {	//函数返回不要用参数返回，一定要用返回值
	listLen := len(list)
	leftList := []byte{}	//初始化切片,无大小指定的无需make
	rightList := []byte{}
	//分组
	for i := 1; i < listLen; i++ {
		if list[0] > list[i] {
			leftList = append(leftList, list[i])
		} else {
			rightList = append(rightList, list[i])
		}
	}

	//是否把子分组也进行同样的算法排列
	isCallSelf := func (listSli []byte) (returnListSli []byte) {	//函数返回不要用参数返回，一定要用返回值，不管是局部函数还是外部函数
		listLen = len(listSli)
		if listLen > 2 {
			returnListSli = QuickSort(listSli)
		} else if listLen == 2 {
			if listSli[0] > listSli[1] {
				listSli[1], listSli[0] = listSli[0], listSli[1]
				returnListSli = listSli
			}
		}
		return
	}
	leftList = isCallSelf(leftList)
	rightList = isCallSelf(rightList)

	//合并
	leftList = append(leftList, list[0])
	list = append(leftList, rightList...)
	returnList = append(leftList, rightList...)
	return
}

func main() {
	list := []byte{1, 43, 54, 62, 21, 66, 100,38, 90, 95, 47, 85, 62, 25, 27, 14, 17, 5, 8, 3, 32, 78, 36, 76}
	v := QuickSort(list)
	fmt.Println("QuickSort =", v)
}
