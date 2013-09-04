package array_util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinimumIsFineWithSingleParam(t *testing.T) {
	assert.Equal(t, Min(1), 1)
}

func TestMinimumIsFineWithTwoParam(t *testing.T) {
	assert.Equal(t, Min(2, 1), 1)
}

func TestMinimumDealsWithSlice(t *testing.T) {
	assert.Equal(t, MinIn([]int{1,2,3}), 1)
}

