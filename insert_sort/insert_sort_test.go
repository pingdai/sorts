package insert_sort

import (
	"fmt"
	"testing"
)

// go test -v --run='InsertSortt'
func TestInsertSortt(t *testing.T) {
	testArr := []int{4, 5, 7, 2, 3}
	insertSort(testArr)
	fmt.Println(testArr)
}
