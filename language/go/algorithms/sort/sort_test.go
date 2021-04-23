package sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	var input = []int{2, 4, 6, 3, 5, 74}
	var expected = []int{2, 3, 4, 5, 6, 74}
	output := Sort(input)
	if !reflect.DeepEqual(expected, output) {
		t.Logf("ERROR: sort of %v should be %v, but get %v", input, expected, output)
		t.Fail()
	}
}
