package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("For a collection of any size", func (t *testing.T) {
		numbers := []int{1,2,3,4,5,6, 7}

		got := Sum(numbers)
	
		want := 28
	
		if got != want {
			t.Errorf("Got %d but received %d given %v", got, want, numbers);
		}
	})
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{3, 6, 1}

	received := SumAll(slice1, slice2)
	expected := []int{ 3, 10}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("Got %v but %v was expected", received, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int {1, 2, 3}, []int {4, 5, 6, 7})
	expected := []int{5, 18}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v but received %v", expected, got)
	}
}