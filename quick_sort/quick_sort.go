package quick_sort

func quickSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}

	// 中间值，默认选取第一个元素
	mid, i := values[0], 1
	// 数组的头尾下标
	head, tail := 0, length-1
	for head < tail {
		if mid < values[i] { // 大于中间值，则把该值挪到中间值的右边
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else { // 小于等于中间值，则把该值挪到中间值的左边
			values[head], values[i] = values[i], values[head]
			i++
			head++
		}
	}

	// 经过上面的for循环后，mid已经被放在了排序的中间位置，再以它为界，进行两段排序
	quickSort(values[:i-1])
	quickSort(values[i:])
}
