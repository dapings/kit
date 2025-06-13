package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	testCases := []struct {
		name string
		data []int
		dst  int
		want int
	}{
		{
			name: "nil slice",
			data: nil,
			dst:  1,
			want: -1,
		},
		{
			name: "empty slice",
			data: []int{},
			dst:  1,
			want: -1,
		},
		{
			name: "exist",
			data: []int{1, 2, 3},
			dst:  2,
			want: 1,
		},
		{
			name: "not exist",
			data: []int{1, 2, 3},
			dst:  4,
			want: -1,
		},
		{
			name: "find first of two elements",
			data: []int{1, 2, 2, 3},
			dst:  2,
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IndexOf(tc.data, tc.dst)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIndexOfByFunc(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	testCases := []struct {
		name  string
		data  []User
		dst   User
		equal equalFunc[User]

		want int
	}{
		{
			name: "nil slice",
			data: nil,
			dst:  User{},
			want: -1,
		},
		{
			name: "empty slice",
			data: []User{},
			dst:  User{},
			want: -1,
		},
		{
			name: "exist",
			data: []User{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
				{Id: 3, Name: "c"},
			},
			dst: User{Id: 2, Name: "b"},
			equal: func(src, dst User) bool {
				return src.Name == dst.Name
			},
			want: 1,
		},
		{
			name: "not exist",
			data: []User{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
				{Id: 3, Name: "c"},
			},
			dst: User{Id: 4, Name: "d"},
			equal: func(src, dst User) bool {
				return src.Name == dst.Name
			},
			want: -1,
		},
		{
			name: "find first of multi elements",
			data: []User{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
				{Id: 3, Name: "d"},
				{Id: 4, Name: "d"},
			},
			dst: User{Id: 3, Name: "d"},
			equal: func(src, dst User) bool {
				return src.Name == dst.Name
			},
			want: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IndexOfByFunc(tc.data, tc.dst, tc.equal)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	testCases := []struct {
		name string
		data []int
		dst  int
		want int
	}{
		{
			name: "nil slice",
			data: nil,
			dst:  1,
			want: -1,
		},
		{
			name: "empty slice",
			data: []int{},
			dst:  1,
			want: -1,
		},
		{
			name: "exist",
			data: []int{1, 2, 3},
			dst:  2,
			want: 1,
		},
		{
			name: "not exist",
			data: []int{1, 2, 3},
			dst:  4,
			want: -1,
		},
		{
			name: "find last of two elements",
			data: []int{1, 2, 2, 3},
			dst:  2,
			want: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := LastIndexOf(tc.data, tc.dst)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLastIndexOfByFunc(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	testCases := []struct {
		name  string
		data  []User
		dst   User
		equal equalFunc[User]

		want int
	}{
		{
			name: "nil slice",
			data: nil,
			dst:  User{},
			want: -1,
		},
		{
			name: "empty slice",
			data: []User{},
			dst:  User{},
			want: -1,
		},
		{
			name: "exist",
			data: []User{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
			},
			dst: User{Id: 2, Name: "b"},
			equal: func(src, dst User) bool {
				return src.Name == dst.Name
			},
			want: 1,
		},
		{
			name: "not exist",
			data: []User{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
			},
			dst: User{Id: 4, Name: "d"},
			equal: func(src, dst User) bool {
				return src.Name == dst.Name
			},
			want: -1,
		},
		{
			name: "find last of multi elements",
			data: []User{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
				{Id: 3, Name: "d"},
				{Id: 4, Name: "d"},
			},
			dst: User{Id: 3, Name: "d"},
			equal: func(src, dst User) bool {
				return src.Name == dst.Name
			},
			want: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := LastIndexOfByFunc(tc.data, tc.dst, tc.equal)
			assert.Equal(t, tc.want, got)
		})
	}
}
