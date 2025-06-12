package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionSet(t *testing.T) {
	testCases := []struct {
		name string
		s1   []int
		s2   []int
		want []int
	}{
		{
			name: "empty slices",
			s1:   []int{},
			s2:   []int{},
			want: []int{},
		},
		{
			name: "empty s2",
			s1:   []int{1, 2, 3},
			s2:   []int{},
			want: []int{},
		},
		{
			name: "not intersection of two slices",
			s1:   []int{1, 2, 3},
			s2:   []int{4, 5, 6},
			want: []int{},
		},
		{
			name: "intersection of two slices",
			s1:   []int{1, 2, 3},
			s2:   []int{3, 4, 5},
			want: []int{3},
		},
		{
			name: "intersection of two slices with duplicates",
			s1:   []int{1, 2, 3, 3},
			s2:   []int{1, 1, 3, 4, 5, 5},
			want: []int{1, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IntersectionSet(tc.s1, tc.s2)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestIntersectionSetByEqFunc(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	testCases := []struct {
		name  string
		s1    []user
		s2    []user
		equal equalFunc[user]
		want  []user
	}{
		{
			name: "two empty slices",
			s1:   []user{},
			s2:   []user{},
			equal: func(u1, u2 user) bool {
				return u1.name == u2.name && u1.age == u2.age
			},
			want: []user{},
		},
		{
			name: "two slices, one of which is an empty slice",
			s1:   []user{{name: "Alice", age: 20}},
			s2:   []user{},
			equal: func(u1, u2 user) bool {
				return u1.name == u2.name && u1.age == u2.age
			},
			want: []user{},
		},
		{
			name: "two slices, non intersection",
			s1:   []user{{name: "Alice", age: 20}, {name: "Bob", age: 30}},
			s2:   []user{{name: "Charlie", age: 40}, {name: "David", age: 50}},
			equal: func(u1, u2 user) bool {
				return u1.name == u2.name && u1.age == u2.age
			},
			want: []user{},
		},
		{
			name: "two slices, intersection",
			s1:   []user{{name: "Alice", age: 20}, {name: "Bob", age: 30}},
			s2:   []user{{name: "Bob", age: 30}, {name: "Charlie", age: 40}},
			equal: func(u1, u2 user) bool {
				return u1.name == u2.name && u1.age == u2.age
			},
			want: []user{{name: "Bob", age: 30}},
		},
		{
			name: "two slices, intersection with duplicates",
			s1:   []user{{name: "Alice", age: 20}, {name: "Bob", age: 30}, {name: "Alice", age: 20}},
			s2:   []user{{name: "Bob", age: 30}, {name: "Charlie", age: 40}, {name: "Bob", age: 30}},
			equal: func(u1, u2 user) bool {
				return u1.name == u2.name && u1.age == u2.age
			},
			want: []user{{name: "Bob", age: 30}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IntersectionSetByEqFunc(tc.s1, tc.s2, tc.equal)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
