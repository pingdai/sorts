package insert_sort

import (
	"fmt"
	"testing"
)

// go test -v --run='InsertSort'
func TestInsertSort(t *testing.T) {
	testArr := []int{4, 5, 7, 2, 3}
	insertSort(testArr)
	fmt.Println(testArr)
}
