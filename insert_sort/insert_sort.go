package insert_sort

func insertSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		// 从主循环的下标开始，依次和左边的元素进行大小比较
		for j := i; j > 0; j-- {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
			}
		}
	}

	return
}
