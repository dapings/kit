package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendDistinct(t *testing.T) {
	testCases := []struct {
		name  string
		src   []int
		items []int
		want  []int
	}{
		{
			name:  "nil src",
			src:   nil,
			items: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "nil items",
			src:   []int{1, 2, 3},
			items: nil,
			want:  []int{1, 2, 3},
		},
		{
			name:  "nil src and items",
			src:   nil,
			items: nil,
			want:  []int{},
		},
		{
			name:  "empty src",
			src:   []int{},
			items: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty items",
			src:   []int{1, 2, 3},
			items: []int{},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty src and items",
			src:   []int{},
			items: []int{},
			want:  []int{},
		},
		{
			name:  "normal src and items, not dup",
			src:   []int{1, 2},
			items: []int{3, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "normal src and items, but contains dup",
			src:   []int{1, 2},
			items: []int{2, 3},
			want:  []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := AppendDistinct(tc.src, tc.items...)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestAddDistinctFunc(t *testing.T) {
	type user struct {
		Name      string
		Telephone string
	}
	testCases := []struct {
		name string

		src   []user
		equal equalFunc[user]
		items []user

		want []user
	}{
		{
			name: "nil src",
			src:  nil,
			equal: func(src, dst user) bool {
				return src.Telephone == dst.Telephone
			},
			items: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "123",
				},
			},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
		},
		{
			name: "nil items",
			src: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "123",
				},
			},
			equal: func(src, dst user) bool {
				return src.Telephone == dst.Telephone
			},
			items: nil,
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
		},
		{
			name: "nil src and items",
			src:  nil,
			equal: func(src, dst user) bool {
				return src.Telephone == dst.Telephone
			},
			items: nil,
			want:  []user{},
		},
		{
			name: "empty src",
			src:  []user{},
			equal: func(src, dst user) bool {
				return src.Telephone == dst.Telephone
			},
			items: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "123",
				},
			},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
		},
		{
			name: "empty items",
			src: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "123",
				},
			},
			equal: func(src, dst user) bool {
				return src.Telephone == dst.Telephone
			},
			items: []user{},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
		},
		{
			name: "empty src and items",
			src:  []user{},
			equal: func(src, dst user) bool {
				return src.Telephone == dst.Telephone
			},
			items: []user{},
			want:  []user{},
		},
		{
			name: "normal src and items, not dup",
			src: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
			equal: func(src, dst user) bool {
				return src.Telephone == dst.Telephone
			},
			items: []user{
				{
					Name:      "u2",
					Telephone: "456",
				},
			},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "456",
				},
			},
		},
		{
			name: "normal src and items, but contains dup",
			src: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "123",
				},
			},
			equal: func(src, dst user) bool {
				return src.Telephone == dst.Telephone
			},
			items: []user{
				{
					Name:      "u2",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "456",
				},
			},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "456",
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := AddDistinctFunc(tc.src, tc.equal, tc.items...)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
