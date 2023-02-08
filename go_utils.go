package go_utils

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func Min[A constraints.Ordered](arr *[]A, size int) (error, any) {

	if size < 1 {
		// var returnSize A = int(size)
		return errors.New("Invalid size array provided"), size
	}

	minEle := (*arr)[0]

	for i := 0; i < size; i += 1 {
		val := (*arr)[i]
		if minEle > val {
			minEle = val
		}
	}

	return nil, minEle
}
