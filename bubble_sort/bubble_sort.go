package bubble_sort

func bubbleSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		// 每一次循环都是把当前循环遍历中找的最大值放在最后面
		for j := 0; j < length-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
}
