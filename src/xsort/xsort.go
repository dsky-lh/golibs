package xsort

// 检测数列是否有序（升序）
func CheckListOrdinal(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i] < list[i-1] {
			return false
		}
	}
	return true
}

// 检测数列是否有序（降序）
func CheckListOrdinalDesc(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			return false
		}
	}
	return true
}

// 简单冒泡排序(升序)
func SimpleBubbleSort(slice []int) {
	length := len(slice)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

// 简单冒泡排序(降序)
func SimpleBubbleSortDesc(slice []int) {
	length := len(slice)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if slice[j] > slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

// 标准冒泡排序(升序)
func StandardBubbleSort(slice []int) {
	length := len(slice)
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

// 标准冒泡排序(降序)
func StandardBubbleSortDesc(slice []int) {
	length := len(slice)
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if slice[j] > slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

// 优化冒泡排序(升序)
func TuningBubbleSort(slice []int) {
	length := len(slice)
	flag := true //加上flag标记，当已经有序时直接结束排序
	for i := 0; i < length && flag; i++ {
		flag = false
		for j := length - 1; j > i; j-- {
			if slice[j] < slice[j-1] {
				flag = true
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

// 优化冒泡排序(降序)
func TuningBubbleSortDesc(slice []int) {
	length := len(slice)
	flag := true //加上flag标记，当已经有序时直接结束排序
	for i := 0; i < length && flag; i++ {
		flag = false
		for j := length - 1; j > i; j-- {
			if slice[j] > slice[j-1] {
				flag = true
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

// 简单选择排序(升序)
func SimpleSelectSort(slice []int) {
	length := len(slice)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		if min != i {
			slice[min], slice[i] = slice[i], slice[min]
		}
	}
}

// 简单选择排序(降序)
func SimpleSelectSortDesc(slice []int) {
	length := len(slice)
	for i := 0; i < length; i++ {
		max := i
		for j := i + 1; j < length; j++ {
			if slice[j] > slice[max] {
				max = j
			}
		}
		if max != i {
			slice[max], slice[i] = slice[i], slice[max]
		}
	}
}

// 直接插入排序(升序)
func DirectInsertSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] { //改为slice[i] > slice[i-1]则是降序
			var j int
			tmp := slice[i]
			for j = i - 1; j >= 0 && tmp < slice[j]; j-- { //改为tmp > slice[j]则是降序
				slice[j+1] = slice[j]
			}
			slice[j+1] = tmp
		}
	}
}

// 直接插入排序(降序)
func DirectInsertSortDesc(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i] > slice[i-1] {
			var j int
			tmp := slice[i]
			for j = i - 1; j >= 0 && tmp > slice[j]; j-- {
				slice[j+1] = slice[j]
			}
			slice[j+1] = tmp
		}
	}
}

// 直接插入排序(升序),此方法时间消耗大概是1的两倍，但是代码简洁一些
func DirectInsertSort2(slice []int) {
	insertSort(slice, 1)
}

// 直接插入排序(升序)，此方法比1时间消耗久一点。
func DirectInsertSort3(slice []int) {
	insertSort2(slice, 1)
}

// 直接插入排序(降序)
func DirectInsertSortDesc2(slice []int) {
	insertSortDesc(slice, 1)
}

// 增量排序(升序)
func insertSort(slice []int, increment int) {
	if increment < 1 { //插入排序最小增量为1
		increment = 1
	}
	for i := increment; i < len(slice); i++ {
		if slice[i] < slice[i-increment] {
			for j := i; j >= increment && slice[j] < slice[j-increment]; j -= increment {
				slice[j], slice[j-increment] = slice[j-increment], slice[j]
			}
		}
	}
}

// 这种方式比insertSort稍微快一些，测过1w条随机数据排序大概是第一种的一倍
// 主要是g第一种a,b = b,a 比第二种多交换了好多次
func insertSort2(slice []int, increment int) {
	if increment < 1 { //插入排序最小增量为1
		increment = 1
	}
	for i := increment; i < len(slice); i++ {
		if slice[i] < slice[i-increment] { //改为slice[i] > slice[i-1]则是降序
			var j int
			tmp := slice[i]
			for j = i - increment; j >= 0 && tmp < slice[j]; j -= increment { //改为tmp > slice[j]则是降序
				slice[j+increment] = slice[j]
			}
			slice[j+increment] = tmp
		}
	}
}

// 增量排序(降序)
func insertSortDesc(slice []int, increment int) {
	if increment < 1 { //插入排序最小增量为1
		increment = 1
	}
	for i := increment; i < len(slice); i++ {
		for j := i; j >= increment && slice[j] > slice[j-increment]; j -= increment {
			slice[j], slice[j-increment] = slice[j-increment], slice[j]
		}
	}
}

// 希尔排序
func ShellSort(slice []int) {
	increment := len(slice)
	for {
		increment = increment/3 + 1
		insertSort2(slice, increment)
		if increment <= 1 { //当增量小于等于1结束排序
			break
		}
	}
}

// 希尔排序
func ShellSortDesc(slice []int) {
	increment := len(slice)
	for {
		increment = increment/3 + 1
		insertSortDesc(slice, increment)
		if increment <= 1 {
			break
		}
	}
}

// 分区，寻找枢轴位置返回
func partition(slice []int, low, high int) int {
	pivotkey := slice[low]
	for low < high {
		for low < high && pivotkey <= slice[high] {
			high--
		}
		slice[low], slice[high] = slice[high], slice[low]
		for low < high && pivotkey >= slice[low] {
			low++
		}
		slice[low], slice[high] = slice[high], slice[low]
	}
	return low
}

// 快速排序递归调用
func quickSort(slice []int, low, high int) {
	if low < high {
		// 寻找枢轴，并递归
		pivot := partition(slice, low, high)
		quickSort(slice, low, pivot-1)
		quickSort(slice, pivot+1, high)
	}
}

// 快速排序
func QuickSort(slice []int) {
	quickSort(slice, 0, len(slice)-1)
}

// 快速排序优化方法：
// 1.三数取中法和九数取中法
// 2.优化不必要的交换
// 3.最小数组时采用直接插入而不采用快速排序
// 4.优化递归操作（采用尾递归的方法）
func partition2(slice []int, low, high int) int {
	pivotkey := slice[low]
	for low < high {
		for low < high && pivotkey <= slice[high] {
			high--
		}
		slice[low] = slice[high]
		for low < high && pivotkey >= slice[low] {
			low++
		}
		slice[high] = slice[low]
	}
	slice[low] = pivotkey
	return low
}

func quickSort2(slice []int, low, high int) {
	for low < high {
		pivot := partition2(slice, low, high)
		quickSort2(slice, low, pivot-1)
		low = pivot + 1
	}
}

func QuickSort2(slice []int) {
	quickSort2(slice, 0, len(slice)-1)
}

// 调整pos位置的结点，使之成为一个大顶堆
func adjustHeap(slice []int, pos int) {
	node := pos
	length := len(slice)
	for node < length/2 {
		// 找到node结点左子结点的位置
		child := 2*node + 1
		// 如果右子结点位置不在数列中且右节点值比左子结点值大，则将child指向右子节点位置
		if child+1 < length && slice[child] < slice[child+1] {
			child++
		}
		// 如果child指向的位置不在数列中或者child位置结点的值不大于父节点
		if child >= length || slice[child] <= slice[node] {
			break
		}
		// 将node结点与其较大子结点的值交换
		slice[node], slice[child] = slice[child], slice[node]
		// 将子结点的位置赋给node继续往下调整
		node = child
	}
}

// 堆排序
func HeapSort(slice []int) {
	// 将数列slice调整为大顶堆,一定得从后往前调整，
	for i := len(slice) - 1; i >= 0; i-- {
		adjustHeap(slice, i)
	}
	for i := len(slice) - 1; i >= 1; i-- {
		// 将大顶堆中根节点值与数列最后一位数交换
		slice[0], slice[i] = slice[i], slice[0]
		// 然后调整根节点位置使剩余数列依然是大顶堆
		adjustHeap(slice[:i-1], 0)
	}
}

func merge(slice []int, low, mid, high int) {
	left := make([]int, 0)
	for i := low; i <= mid; i++ {
		left = append(left, slice[i])
	}
	right := make([]int, 0)
	for j := mid + 1; j <= high; j++ {
		right = append(right, slice[j])
	}

	i, j := 0, 0
	leftLen, rightLen := len(left), len(right)
	for k := low; k <= high; k++ {
		// 如果左集合已经遍历完
		if i >= leftLen {
			slice[k] = right[j]
			j++
			continue
		}
		// 如果右集合已经遍历完
		if j >= rightLen {
			slice[k] = left[i]
			i++
			continue
		}
		// 左右集合都没有便利完
		if left[i] <= right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
}

func mergeSort(slice []int, low, high int) {
	if low < high {
		mid := (high + low) / 2
		mergeSort(slice, low, mid)
		mergeSort(slice, mid+1, high)
		merge(slice, low, mid, high)
	}
}

// 归并排序
func MergeSort(slice []int) {
	mergeSort(slice, 0, len(slice)-1)
}
