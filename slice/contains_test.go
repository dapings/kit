package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		data []T
		item T
		want bool
	}
	tests := []testCase[int]{
		{
			name: "nil slice",
			item: 2,
			want: false,
		},
		{
			name: "empty slice",
			data: []int{},
			item: 2,
			want: false,
		},
		{
			name: "not existed element in slice",
			data: []int{1, 3, 5, 7, 9},
			item: 2,
			want: false,
		},
		{
			name: "existed element in slice",
			data: []int{1, 3, 5, 7, 9},
			item: 5,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Contains(tt.data, tt.item))
		})
	}
}

func TestContainsByFunc(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		data []T
		dst  T
		want bool
	}
	tests := []testCase[int]{
		{
			name: "nil slice",
			dst:  2,
			want: false,
		},
		{
			name: "empty slice",
			data: []int{},
			dst:  2,
			want: false,
		},
		{
			name: "not existed element in slice",
			data: []int{1, 3, 5, 7, 9},
			dst:  2,
			want: false,
		},
		{
			name: "existed element in slice",
			data: []int{1, 3, 5, 7, 9},
			dst:  5,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ContainsByFunc(tt.data, tt.dst, func(s, d int) bool {
				return s == d
			}))
		})
	}
}
