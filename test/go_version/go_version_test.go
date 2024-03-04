package go_version_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForLoopRangeScope(t *testing.T) {
	// This test should pass with the full Go 1.22 compilation experience
	var wg sync.WaitGroup
	wg.Add(10)
	s := make([]int, 10)
	for i := range 10 {
		go func() {
			s[i] = i
			wg.Done()
		}()
	}
	wg.Wait()
	assert.EqualValues(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s)
}
