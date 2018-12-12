package binary_sort

import (
	"fmt"
	"testing"
)

// go test -v --run='BinarySort'
func TestBinarySort(t *testing.T) {
	testArr := []int{4, 5, 7, 2, 3}
	binarySort(testArr)
	fmt.Println(testArr)
}
