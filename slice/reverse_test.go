package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name string
		data []int

		want []int
	}{
		{
			name: "nil slice",
			data: nil,
			want: []int{},
		},
		{
			name: "empty slice",
			data: []int{},
			want: []int{},
		},
		{
			name: "slice with one element",
			data: []int{1},
			want: []int{1},
		},
		{
			name: "slice with multiple elements",
			data: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Reverse(tc.data)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestReverseInplace(t *testing.T) {
	testCases := []struct {
		name string
		data []int
		want []int
	}{
		{
			name: "nil slice",
			data: nil,
			want: nil,
		},
		{
			name: "empty slice",
			data: []int{},
			want: []int{},
		},
		{
			name: "slice with one element",
			data: []int{1},
			want: []int{1},
		},
		{
			name: "slice with multiple elements",
			data: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ReverseInplace(tc.data)
			assert.Equal(t, tc.want, tc.data)
		})
	}
}
