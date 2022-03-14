package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortEmptyArray(t *testing.T) {
	arg := []int{}
	gotVal := quicksort(arg)

	wantVal := []int{}

	assert.Equal(t, wantVal, gotVal)
}

func TestQuickSortDifferentSizeFalse(t *testing.T) {
	arg := []int{1, 5, 2}
	gotVal := quicksort(arg)

	wantVal := []int{1, 2, 3, 5}

	assert.NotEqual(t, wantVal, gotVal)
}

func TestQuickSortRepeatedEvenElementsFalse(t *testing.T) {
	arg := []int{1, 5, 2, 5}
	gotVal := quicksort(arg)

	wantVal := []int{1, 2, 5}

	assert.NotEqual(t, wantVal, gotVal)
}

func TestQuickSortRepeatedEvenElementsTrue(t *testing.T) {
	arg := []int{1, 5, 2, 5}
	gotVal := quicksort(arg)

	wantVal := []int{1, 2, 5, 5}

	assert.Equal(t, wantVal, gotVal)
}

func TestQuickSortRepeatedOddElementsFalse(t *testing.T) {
	arg := []int{1, 5, 2, 5, 0}
	gotVal := quicksort(arg)

	wantVal := []int{0, 1, 2, 5}

	assert.NotEqual(t, wantVal, gotVal)
}

func TestQuickSortRepeatedOddElementsTrue(t *testing.T) {
	arg := []int{1, 5, 2, 5, 0}
	gotVal := quicksort(arg)

	wantVal := []int{0, 1, 2, 5, 5}

	assert.Equal(t, wantVal, gotVal)
}
