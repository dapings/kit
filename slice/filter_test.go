package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name   string
		data   []int
		filter filterFunc[int]
		want   []int
	}{
		{
			name:   "odd number",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			filter: func(idx int, item int) bool { return item%2 == 1 },
			want:   []int{1, 3, 5, 7, 9},
		},
		{
			name:   "even number",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			filter: func(idx int, item int) bool { return item%2 == 0 },
			want:   []int{2, 4, 6, 8, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Filter(tc.data, tc.filter)
			assert.Equal(t, tc.want, got)
		})
	}
}
