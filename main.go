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

//简单冒泡排序
func TestBubbleSort(num int) {
	fmt.Println("simple bubble sort with num of:", num)
	list := genIntSlice(num, num)
	start := time.Now().UnixNano()
	xsort.SimpleBubbleSort(list)
	end := time.Now().UnixNano()
	fmt.Println("simple bubble sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
}

//标准冒泡排序
func TestBubbleSort2(num int) {
	fmt.Println("standard bubble with num of:", num)
	list := genIntSlice(num, num)
	start := time.Now().UnixNano()
	xsort.StandardBubbleSort(list)
	end := time.Now().UnixNano()
	fmt.Println("standard bubble sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
}

//优化冒泡排序
func TestBubbleSort3(num int) {
	fmt.Println("tuning bubble sort with num of:", num)
	list := genIntSlice(num, num)
	start := time.Now().UnixNano()
	xsort.TuningBubbleSort(list)
	end := time.Now().UnixNano()
	fmt.Println("tuning bubble sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
}

//简单选择排序
func TestSelectSort(num int) {
	fmt.Println("simple select sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.SimpleSelectSort(list)
	end := time.Now().UnixNano()
	fmt.Println("simple select sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//直接插入排序
func TestInsertSort(num int) {
	fmt.Println("direct insert sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.DirectInsertSort(list)
	end := time.Now().UnixNano()
	fmt.Println("direct insert sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//复用希尔排序中比较函数的插入排序，性能比第一种差一些
func TestInsertSort2(num int) {
	fmt.Println("direct insert sort 2 with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.DirectInsertSort2(list)
	end := time.Now().UnixNano()
	fmt.Println("direct insert sort 2 use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//对第二种做了优化，没有使用go语言特性，性能比第二种强些
func TestInsertSort3(num int) {
	fmt.Println("direct insert sort 3 with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.DirectInsertSort3(list)
	end := time.Now().UnixNano()
	fmt.Println("direct insert sort 3 use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//希尔排序
func TestShellSort(num int) {
	fmt.Println("shell sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.ShellSort(list)
	// xsort.ShellSortDesc(list)
	end := time.Now().UnixNano()
	fmt.Println("shell sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//快速排序
func TestQuickSort(num int) {
	fmt.Println("quick sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.QuickSort(list)
	end := time.Now().UnixNano()
	fmt.Println("quick sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//快速排序，对第一种做了部分优化，整体差别不大，因为快速排序本来就很快了
func TestQuickSort2(num int) {
	fmt.Println("quick sort 2 with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.QuickSort2(list)
	end := time.Now().UnixNano()
	fmt.Println("quick sort 2 use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//堆排序
func TestHeapSort(num int) {
	fmt.Println("heap sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.HeapSort(list)
	end := time.Now().UnixNano()
	fmt.Println("heap sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//归并排序
func TestMergeSort(num int) {
	fmt.Println("merge sort with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.MergeSort(list)
	end := time.Now().UnixNano()
	fmt.Println("merge sort use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//第二种归并排序,实现方法有些差异，利用了go的一些特性
func TestMergeSort2(num int) {
	fmt.Println("merge sort 2 with num of:", num)
	list := genIntSlice(num, num)
	// fmt.Println("list:", list)
	start := time.Now().UnixNano()
	xsort.MergeSort2(&list)
	end := time.Now().UnixNano()
	fmt.Println("merge sort 2 use time:", (end-start)/1e6, "ms")
	if !xsort.CheckListOrdinal(list) {
		fmt.Println("sort err")
		return
	}
	// fmt.Println("sort list:", list)
}

//测试十万条数据，七种排序方法所需要的时间
func TestSevenSort_100000() {
	fmt.Println("simple sort....")
	// 冒泡排序
	TestBubbleSort(100000)
	TestBubbleSort2(100000)
	TestBubbleSort3(100000)
	// 选择排序
	TestSelectSort(100000)
	// 插入排序
	TestInsertSort(100000)
	TestInsertSort2(100000)
	TestInsertSort3(100000)

	fmt.Println("\nadvanced sort....")
	// 希尔排序
	TestShellSort(100000)
	// 堆排序
	TestHeapSort(100000)
	// 快速排序
	TestQuickSort(100000)
	TestQuickSort2(100000)
	// 归并排序
	TestMergeSort(100000)
	TestMergeSort2(100000)
}

func TestAdvancedSort_10000000() {
	fmt.Println("\nadvanced sort big data....")
	// 希尔排序
	TestShellSort(10000000)
	// 堆排序
	TestHeapSort(10000000)
	// 快速排序
	TestQuickSort(10000000)
	TestQuickSort2(10000000)
	// 归并排序
	TestMergeSort(10000000)
	TestMergeSort2(10000000)
}

func main() {
	TestSevenSort_100000()
	TestAdvancedSort_10000000()
}
