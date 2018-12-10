package quick_sort

import (
	"fmt"
	"testing"
)

// go test -v --run='QuickSort'
func TestQuickSort(t *testing.T) {
	testArr := []int{4, 5, 7, 2, 3}
	quickSort(testArr)
	fmt.Println(testArr)
}
