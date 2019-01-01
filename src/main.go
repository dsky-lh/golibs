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
	xsort.SimpleBubbleSort(list)
	end = time.Now().UnixNano()
	fmt.Println("simple bubble sort use ", (end-start)/1e6, "ms")

	list = genIntSlice(num, num)
	start = time.Now().UnixNano()
	xsort.StandardBubbleSort(list)
	end = time.Now().UnixNano()
	fmt.Println("standard bubble sort use ", (end-start)/1e6, "ms")

	list = genIntSlice(num, num)
	start = time.Now().UnixNano()
	xsort.TuningBubbleSort(list)
	end = time.Now().UnixNano()
	fmt.Println("tuning bubble sort use ", (end-start)/1e6, "ms")
}

func TestSelectSort(num int) {
	fmt.Println("test select sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.SimpleSelectSort(list)
	end := time.Now().UnixNano()
	fmt.Println("simple select sort use ", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

func TestInsertSort(num int) {
	fmt.Println("test insert sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.DirectInsertSort(list)
	end := time.Now().UnixNano()
	fmt.Println("direct insert sort use ", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

func TestInsertSort2(num int) {
	fmt.Println("test insert sort 2 with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.DirectInsertSort2(list)
	end := time.Now().UnixNano()
	fmt.Println("direct insert sort 2 use ", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

func TestInsertSort3(num int) {
	fmt.Println("test insert sort 3 with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.DirectInsertSort3(list)
	end := time.Now().UnixNano()
	fmt.Println("direct insert sort 3 use ", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

func TestInsertSortDesc(num int) {
	fmt.Println("test insert sort desc with num of:", num)
	list := genIntSlice(num, num)
	fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.DirectInsertSortDesc(list)
	end := time.Now().UnixNano()
	fmt.Println("direct insert sort use ", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinalDesc(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

func TestShellSort(num int) {
	fmt.Println("test shell sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.ShellSort(list)
	// xsort.ShellSortDesc(list)
	end := time.Now().UnixNano()
	fmt.Println("shell sort use ", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

func TestShellSortDesc(num int) {
	fmt.Println("test shell sort desc with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.ShellSortDesc(list)
	end := time.Now().UnixNano()
	fmt.Println("shell sort desc use ", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinalDesc(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

func TestQuickSort(num int) {
	fmt.Println("test quick sort desc with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.QuickSort(list)
	end := time.Now().UnixNano()
	fmt.Println("quick sort desc use ", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

func main() {
	// TestBubbleSort(1000)
	// TestBubbleSort(10000)
	// TestBubbleSort(100000)
	// // TestBubbleSort(1000000)

	// TestSelectSort(100)
	// TestSelectSort(10000)
	// TestSelectSort(100000)

	// TestInsertSort(100000)
	// TestInsertSort2(100000)
	// TestInsertSort3(100000)
	// TestInsertSort(10000)
	// TestInsertSort(100000)
	// TestInsertSortDesc(100)

	TestShellSort(10000000)
	// TestShellSortDesc(10000000)

	TestQuickSort(10000000)
}
