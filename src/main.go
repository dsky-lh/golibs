package main

import (
	"fmt"
	"math/rand"
	"time"
	"xsort"
)

// 随机生成固定范围内固定数量的int类型切片
func genIntSlice(maxNum, size int) []int {
	intSlice := make([]int, 0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		num := rand.Intn(maxNum)
		intSlice = append(intSlice, num)
	}
	return intSlice
}

//测试指定大小数组的排序性能
func TestBubbleSort(num int) {
	var list []int
	var start, end int64
	fmt.Println("test bubble sort with num of:", num)
	list = genIntSlice(num, num)
	start = time.Now().UnixNano()
	xsort.SimpleBubbleSort(list, len(list))
	end = time.Now().UnixNano()
	fmt.Println("simple bubble sort use ", (end-start)/1e6, "ms")

	list = genIntSlice(num, num)
	start = time.Now().UnixNano()
	xsort.StandardBubbleSort(list, len(list))
	end = time.Now().UnixNano()
	fmt.Println("standard bubble sort use ", (end-start)/1e6, "ms")

	list = genIntSlice(num, num)
	start = time.Now().UnixNano()
	xsort.TuningBubbleSort(list, len(list))
	end = time.Now().UnixNano()
	fmt.Println("tuning bubble sort use ", (end-start)/1e6, "ms")
}

func TestSelectSort(num int) {
	fmt.Println("test select sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.SimpleSelectSort(list, len(list))
	end := time.Now().UnixNano()
	fmt.Println("simple select sort use ", (end-start)/1e6, "ms")
	// fmt.Println("sort list:", list)
}

func TestInsertSort(num int) {
	fmt.Println("test insert sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.DirectInsertSort(list, len(list))
	end := time.Now().UnixNano()
	fmt.Println("direct insert sort use ", (end-start)/1e6, "ms")
	// fmt.Println("sort list:", list)
}

func main() {
	TestBubbleSort(1000)
	TestBubbleSort(10000)
	TestBubbleSort(100000)
	// TestBubbleSort(1000000)

	TestSelectSort(100)
	TestSelectSort(10000)
	TestSelectSort(100000)

	TestInsertSort(1000)
	TestInsertSort(10000)
	TestInsertSort(100000)
}
