package xsort

// 简单冒泡排序(升序)
func SimpleBubbleSort(slice []int, length int) {
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

// 简单冒泡排序(降序)
func SimpleBubbleSortDesc(slice []int, length int) {
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if slice[j] > slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

// 标准冒泡排序(升序)
func StandardBubbleSort(slice []int, length int) {
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

// 标准冒泡排序(降序)
func StandardBubbleSortDesc(slice []int, length int) {
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if slice[j] > slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

// 优化冒泡排序(升序)
func TuningBubbleSort(slice []int, length int) {
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
func TuningBubbleSortDesc(slice []int, length int) {
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
func SimpleSelectSort(slice []int, length int) {
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
func SimpleSelectSortDesc(slice []int, length int) {
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
func DirectInsertSort(slice []int, length int) {
	for i := 1; i < length; i++ {
		if slice[i] < slice[i-1] {
			var j int
			tmp := slice[i]
			for j = i - 1; j >= 0 && tmp < slice[j]; j-- {
				slice[j+1] = slice[j]
			}
			slice[j+1] = tmp
		}
	}
}

// 直接插入排序(降序)
func DirectInsertSortDesc(slice []int, length int) {
	for i := 1; i < length; i++ {
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
