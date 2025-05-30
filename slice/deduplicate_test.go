package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicate(t *testing.T) {
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
			name: "not duplicated element slice",
			data: []int{1, 3, 5, 7, 9},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			name: "duplicated element slice",
			data: []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 9},
			want: []int{1, 3, 5, 7, 9},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, Deduplicate(tc.data))
		})
	}
}

func TestDeduplicateByEqFunc(t *testing.T) {
	type User struct {
		Id   int
		Name string
		Age  int
	}
	testCases := []struct {
		name   string
		data   []User
		eqFunc equalFunc[User]
		want   []User
	}{
		{
			name: "nil slice",
			data: nil,
			want: []User{},
		},
		{
			name: "empty slice",
			data: []User{},
			want: []User{},
		},
		{
			name: "remove the User.Name duplicated element of slice, but slice unchanged",
			data: []User{
				{Id: 1, Name: "Tom", Age: 18},
				{Id: 2, Name: "Jack", Age: 20},
			},
			eqFunc: func(src, dst User) bool {
				return src.Name == dst.Name
			},
			want: []User{
				{Id: 1, Name: "Tom", Age: 18},
				{Id: 2, Name: "Jack", Age: 20},
			},
		},
		{
			name: "remove the User.Name duplicated element of slice, slice changed",
			data: []User{
				{Id: 1, Name: "Tom", Age: 18},
				{Id: 2, Name: "Jack", Age: 20},
				{Id: 3, Name: "Tom", Age: 19},
			},
			eqFunc: func(src, dst User) bool {
				return src.Name == dst.Name
			},
			want: []User{
				{Id: 1, Name: "Tom", Age: 18},
				{Id: 2, Name: "Jack", Age: 20},
			},
		},
		{
			name: "remove the User.Age duplicated element of slice, slice changed",
			data: []User{
				{Id: 1, Name: "Tom", Age: 18},
				{Id: 2, Name: "Jack", Age: 20},
				{Id: 3, Name: "Alert", Age: 18},
			},
			eqFunc: func(src, dst User) bool {
				return src.Age == dst.Age
			},
			want: []User{
				{Id: 1, Name: "Tom", Age: 18},
				{Id: 2, Name: "Jack", Age: 20},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, DeduplicateByEqFunc(tc.data, tc.eqFunc))
		})
	}
}
