package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrop(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	stream := Drop[int](5)(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{6, 7, 8, 9, 10}, result)
}

func TestDropAll(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4})
	stream := Drop[int](5)(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Empty(t, result)
}
