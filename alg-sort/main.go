package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	num      = 10000  // 测试数组的长度
	rangeNum = 100000 // 数组元素大小范围
)

func main() {
	arr := GenerateRand() //生成随机数组
	//排序前 复制原数组
	org_arr := make([]int, num)
	copy(org_arr, arr)
	Bubble(arr)
	sort.Ints(org_arr) //使sort模块对原数组排序
	fmt.Println(arr[:15])
	fmt.Println(org_arr[:15])
	fmt.Println(IsSame(arr, org_arr))
}

func GenerateRand() []int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = randSeed.Intn(rangeNum)
	}
	return arr
}

//比较两个切片
func IsSame(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	if (slice1 == nil) != (slice2 == nil) {
		return false
	}

	for i, num := range slice1 {
		if num != slice2[i] {
			return false
		}
	}
	return true
}

//冒泡排序
func Bubble(arr []int) {
	size := len(arr)
	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped != true {
			break
		}
	}
}
