package go_version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForLoopRangeScope(t *testing.T) {
	// This test should pass with the full Go 1.22 compilation experience
	s := make([]int, 10)
	for i := range 10 {
		go func() {
			s[i] = i
		}()
	}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s)
}
