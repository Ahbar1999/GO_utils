package go_utils

import "testing"

func TestMinOnIntegers(t *testing.T) {
	nums := []int{12, 4, 1, 51, 0, -11}
	want := -11
	err, ans := Min(&nums, len(nums))
	if err != nil && ans != want {
		t.Fatalf(`Call to Min() failed for nums: %v`, nums)
	}
}

func TestMinOnBytes(t *testing.T) {
	chars := []byte{'b', 'd', 'e', 'A', 'B', 'a'}
	want := byte('A')
	err, ans := Min(&chars, len(chars))
	if err != nil && ans != want {
		t.Fatalf(`Call to Min() failed for nums: %v`, chars)
	}
}

func TestMinOnEmptyArray(t *testing.T) {
	chars := []int{}
	err, size := Min(&chars, len(chars))
	if err != nil && size != len(chars) {
		t.Fatalf(`Call to Min() failed for chars: %v`, chars)
	}
}
