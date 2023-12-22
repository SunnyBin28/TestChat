// main.go

package main

import (
	"fmt"
)

// 冒泡排序算法实现
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// 示例数据
	data := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("未排序的数组:", data)

	// 调用冒泡排序算法
	bubbleSort(data)

	fmt.Println("排序后的数组:", data)
}
