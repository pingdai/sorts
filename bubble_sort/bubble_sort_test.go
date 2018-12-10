package bubble_sort

import (
	"fmt"
	"testing"
)

// go test -v --run='BubbleSort'
func TestBubbleSort(t *testing.T) {
	testArr := []int{4, 5, 7, 2, 3}
	bubbleSort(testArr)
	fmt.Println(testArr)
}
