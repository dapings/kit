package slice

import (
	"testing"

	"github.com/dapings/kit/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestDeleteByFilterFunc(t *testing.T) {
	type testCase[T any] struct {
		name   string
		data   []T
		filter filterFunc[T]

		want []T
	}
	testCases := []testCase[int]{
		{
			name:   "empty slice",
			data:   []int{},
			filter: func(idx int, src int) bool { return false },
			want:   []int{},
		},
		{
			name:   "not del element",
			data:   []int{1, 2, 3, 4, 5},
			filter: func(idx int, src int) bool { return false },
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "del first element",
			data:   []int{0, 1, 2, 3, 4, 5},
			filter: func(idx int, src int) bool { return idx == 0 },
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "del the first two elements",
			data:   []int{0, 1, 2, 3, 4, 5},
			filter: func(idx int, src int) bool { return idx < 2 },
			want:   []int{2, 3, 4, 5},
		},
		{
			name:   "del a element in the middle",
			data:   []int{0, 1, 2, 3, 4, 5, 6},
			filter: func(idx int, src int) bool { return idx == 3 },
			want:   []int{0, 1, 2, 4, 5, 6},
		},
		{
			name: "del multi discrete elements in the middle",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filter: func(idx int, src int) bool {
				return idx == 3 || idx == 5
			},
			want: []int{0, 1, 2, 4, 6, 7},
		},
		{
			name: "del multi consecutive elements in the middle",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filter: func(idx int, src int) bool {
				return idx == 3 || idx == 4
			},
			want: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name: "del multi elements in the middle. The first part is one element, and the second part is the discrete element",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filter: func(idx int, src int) bool {
				return idx == 2 || idx == 4 || idx == 5
			},
			want: []int{0, 1, 3, 6, 7},
		},
		{
			name: "del multi elements in the middle. The first part is the discrete element, and the second part is one element",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filter: func(idx int, src int) bool {
				return idx == 2 || idx == 3 || idx == 5
			},
			want: []int{0, 1, 4, 6, 7},
		},
		{
			name: "del last two elements",
			data: []int{0, 1, 2, 3, 4, 5, 6},
			filter: func(idx int, src int) bool {
				return idx == 5 || idx == 6
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "del last element",
			data: []int{0, 1, 2, 3, 4, 5, 6},
			filter: func(idx int, src int) bool {
				return idx == 6
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "del all elements",
			data: []int{0, 1, 2, 3, 4, 5, 6},
			filter: func(idx int, src int) bool {
				return true
			},
			want: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteByFilterFunc(tc.data, tc.filter)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDeleteByIndex(t *testing.T) {
	type testCase[T any] struct {
		name    string
		data    []T
		index   int
		want    []T
		wantErr error
	}
	testCases := []testCase[int]{
		{
			name:    "index out of range, index less than 0",
			data:    []int{1, 2, 3, 4, 5},
			index:   -1,
			want:    nil,
			wantErr: errors.NewIndexOutOfRange(5, -1),
		},
		{
			name:    "index out of range, index equal length",
			data:    []int{1, 2, 3, 4, 5},
			index:   5,
			want:    nil,
			wantErr: errors.NewIndexOutOfRange(5, 5),
		},
		{
			name:    "empty slice",
			data:    []int{},
			index:   0,
			want:    nil,
			wantErr: errors.NewIndexOutOfRange(0, 0),
		},
		{
			name:    "del the unique element",
			data:    []int{1},
			index:   0,
			want:    []int{},
			wantErr: nil,
		},
		{
			name:    "del the first element",
			data:    []int{1, 2, 3},
			index:   0,
			want:    []int{2, 3},
			wantErr: nil,
		},
		{
			name:    "del the element of index 2",
			data:    []int{1, 2, 3},
			index:   2,
			want:    []int{1, 2},
			wantErr: nil,
		},
		{
			name:    "del the element of index 1",
			data:    []int{1, 2, 3},
			index:   1,
			want:    []int{1, 3},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := DeleteByIndex(tc.data, tc.index)
			if tc.wantErr != nil || err != nil {
				assert.Equal(t, tc.wantErr, err)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDeleteByItem(t *testing.T) {
	testCases := []struct {
		name string
		data []int
		item int
		want []int
	}{
		{
			name: "empty slice",
			data: []int{},
			item: 1,
			want: []int{},
		},
		{
			name: "del the non-existed element",
			data: []int{2, 4, 6, 8},
			item: 1,
			want: []int{2, 4, 6, 8},
		},
		{
			name: "del the first element",
			data: []int{2, 4, 6, 8},
			item: 2,
			want: []int{4, 6, 8},
		},
		{
			name: "del the last element",
			data: []int{2, 4, 6, 8},
			item: 8,
			want: []int{2, 4, 6},
		},
		{
			name: "del the multi element",
			data: []int{2, 4, 6, 4, 8, 4},
			item: 4,
			want: []int{2, 6, 8},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteByItem(tc.data, tc.item)
			assert.Equal(t, tc.want, got)
		})
	}
}
