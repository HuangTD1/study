package main

import (
	"fmt"
	"sort"
)

/*
原二维数组
1	9	5    -----4
2	3	6    -----1
3	6	9    -----2
1	8	3    -----3

排序后
2	3	6
3	6	9
1	8	3
1	9	5

*/

func main() {
	// 定义二维数组
	nums := [][]int{{1, 9, 5}, {2, 3, 6}, {3, 6, 9}, {1, 8, 3}}
	firstIndex := 2 //第二排排序
	result := ArraySort(nums, firstIndex-1)
	fmt.Println(result)
}

// 按指定规则对numArray 进行排序
func ArraySort(numArray [][]int, firstIndex int) [][]int {
	// 检查
	if len(numArray) <= 1 {
		return numArray
	}
	if firstIndex < 0 || firstIndex > len(numArray[0])-1 {
		fmt.Println("Warning: Param firstIndex should between 0 and len(numArray)-1.The original array is returned.")
		return numArray
	}
	// 排序
	mIntArray := &IntArray{numArray, firstIndex}
	sort.Sort(mIntArray)
	return mIntArray.mArr
}

type IntArray struct {
	mArr       [][]int
	firstIndex int
}

// IntArray 实现sort.Interface 接口
func (arr *IntArray) Len() int {
	return len(arr.mArr)
}

func (arr *IntArray) Swap(i, j int) {
	arr.mArr[i], arr.mArr[j] = arr.mArr[j], arr.mArr[i]
}

func (arr *IntArray) Less(i, j int) bool {
	arr1 := arr.mArr[i]
	arr2 := arr.mArr[j]
	for index := arr.firstIndex; index < len(arr1); index++ {
		if arr1[index] < arr2[index] {
			return true
		} else if arr1[index] > arr2[index] {
			return false
		}
	}
	return i < j
}
