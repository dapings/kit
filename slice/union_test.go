package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	testCases := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{
			name:  "two nil slices",
			input: [][]int{nil, nil},
			want:  []int{},
		},
		{
			name:  "two empty slices",
			input: [][]int{{}, {}},
			want:  []int{},
		},
		{
			name:  "two slices with same elements, where one slice contains a dup element",
			input: [][]int{{1, 2, 3, 4}, {1, 2, 2, 3, 5}},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "two slices with different elements",
			input: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:  "three nil slices",
			input: [][]int{nil, nil, nil},
			want:  []int{},
		},
		{
			name:  "three empty slices",
			input: [][]int{{}, {}, {}},
			want:  []int{},
		},
		{
			name:  "three slices with same elements, where two slice contains a dup element",
			input: [][]int{{1, 2, 3, 4}, {1, 2, 2, 3, 5}, {5, 5, 6, 7}},
			want:  []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:  "three slices with different elements",
			input: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Union(tc.input...)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestUnionByEqFunc(t *testing.T) {
	type user struct {
		id   int
		name string
	}
	testCases := []struct {
		name  string
		input [][]user
		equal equalFunc[user]
		want  []user
	}{
		{
			name:  "two nil slices",
			input: [][]user{nil, nil},
			equal: func(u1, u2 user) bool {
				return u1.id == u2.id
			},
			want: []user{},
		},
		{
			name:  "two empty slices",
			input: [][]user{{}, {}},
			equal: func(u1, u2 user) bool {
				return u1.id == u2.id
			},
			want: []user{},
		},
		{
			name: "two slices with same elements, where one slice contains a dup element",
			input: [][]user{
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}},
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 5, name: "5"}},
			},
			equal: func(u1, u2 user) bool {
				return u1.id == u2.id
			},
			want: []user{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}, {id: 5, name: "5"}},
		},
		{
			name: "two slices with different elements",
			input: [][]user{
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}},
				{{id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}},
			},
			equal: func(u1, u2 user) bool {
				return u1.id == u2.id
			},
			want: []user{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}, {id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}},
		},
		{
			name:  "three nil slices",
			input: [][]user{nil, nil, nil},
			equal: func(u1, u2 user) bool {
				return u1.id == u2.id
			},
			want: []user{},
		},
		{
			name:  "three empty slices",
			input: [][]user{{}, {}, {}},
			equal: func(u1, u2 user) bool {
				return u1.id == u2.id
			},
			want: []user{},
		},
		{
			name: "three slices with same elements, where one slice contains a dup element",
			input: [][]user{
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}},
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 5, name: "5"}},
				{{id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}, {id: 9, name: "9"}},
			},
			equal: func(u1, u2 user) bool {
				return u1.id == u2.id
			},
			want: []user{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}, {id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}, {id: 9, name: "9"}},
		},
		{
			name: "three slices with different elements",
			input: [][]user{
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}},
				{{id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}},
				{{id: 9, name: "9"}, {id: 10, name: "10"}},
			},
			equal: func(u1, u2 user) bool {
				return u1.id == u2.id
			},
			want: []user{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}, {id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}, {id: 9, name: "9"}, {id: 10, name: "10"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := UnionByEqFunc(tc.equal, tc.input...)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
