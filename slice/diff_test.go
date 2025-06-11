package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	testCases := []struct {
		name string
		s1   []int
		s2   []int
		want []int
	}{
		{
			name: "nil slice",
			s1:   nil,
			s2:   nil,
			want: []int{},
		},
		{
			name: "empty slice",
			s1:   []int{},
			s2:   []int{},
			want: []int{},
		},
		{
			name: "s1 is nil",
			s1:   nil,
			s2:   []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "s2 is nil",
			s1:   []int{1, 2, 3},
			s2:   nil,
			want: []int{1, 2, 3},
		},
		{
			name: "s1 is empty",
			s1:   []int{},
			s2:   []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "s2 is empty",
			s1:   []int{1, 2, 3},
			s2:   []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "s1 and s2 are equal",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "s1 is subset of s2",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2, 3, 4, 5},
			want: []int{},
		},
		{
			name: "s2 is subset of s1",
			s1:   []int{1, 2, 3, 4, 5},
			s2:   []int{1, 2, 3},
			want: []int{4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Diff(tc.s1, tc.s2)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestDiffFunc(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	testCases := []struct {
		name  string
		s1    []User
		s2    []User
		equal equalFunc[User]
		want  []User
	}{
		{
			name:  "nil slice",
			s1:    nil,
			s2:    nil,
			equal: nil,
			want:  []User{},
		},
		{
			name:  "empty slice",
			s1:    []User{},
			s2:    []User{},
			equal: nil,
			want:  []User{},
		},
		{
			name: "s1 and equal is nil",
			s1:   nil,
			s2: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			equal: nil,
			want:  []User{},
		},
		{
			name: "s2 and equal is nil",
			s1: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			s2:    nil,
			equal: nil,
			want: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
		},
		{
			name: "s1 is nil",
			s1:   nil,
			s2: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			equal: func(s1, s2 User) bool {
				return s1.Name == s2.Name
			},
			want: []User{},
		},
		{
			name: "s2 is nil",
			s1: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			s2: nil,
			equal: func(s1, s2 User) bool {
				return s1.Name == s2.Name
			},
			want: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
		},
		{
			name: "s1 is empty",
			s1:   []User{},
			s2: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			equal: nil,
			want:  []User{},
		},
		{
			name: "s2 is empty",
			s1: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			s2: []User{},
			equal: func(s1, s2 User) bool {
				return s1.Name == s2.Name
			},
			want: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
		},
		{
			name: "s1 and s2 are equal",
			s1: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			s2: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			equal: func(s1, s2 User) bool {
				return s1.Name == s2.Name
			},
			want: []User{},
		},
		{
			name: "s1 is subset of s2",
			s1: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			s2: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
				{Id: 4, Name: "Joe"},
				{Id: 5, Name: "Jill"},
			},
			equal: func(s1, s2 User) bool {
				return s1.Name == s2.Name
			},
			want: []User{},
		},
		{
			name: "s2 is subset of s1",
			s1: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
				{Id: 4, Name: "Joe"},
				{Id: 5, Name: "Jill"},
			},
			s2: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Jane"},
				{Id: 3, Name: "Jim"},
			},
			equal: func(s1, s2 User) bool {
				return s1.Name == s2.Name
			},
			want: []User{
				{Id: 4, Name: "Joe"},
				{Id: 5, Name: "Jill"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := DiffFunc(tc.s1, tc.s2, tc.equal)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
