package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toMap(t *testing.T) {
	testCases := []struct {
		name string
		data []int
		want map[int]struct{}
	}{
		{
			name: "nil slice",
			data: nil,
			want: map[int]struct{}{},
		},
		{
			name: "empty slice",
			data: []int{},
			want: map[int]struct{}{},
		},
		{
			name: "slice with the duplicated elements",
			data: []int{1, 2, 2, 3},
			want: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
		{
			name: "slice without the duplicated elements",
			data: []int{1, 2, 3},
			want: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, toMap(tc.data))
		})
	}
}

func Test_mapKeyToSlice(t *testing.T) {
	testCases := []struct {
		name string
		data map[int]struct{}
		want []int
	}{
		{
			name: "make the nil map to slice",
			data: nil,
			want: []int{},
		},
		{
			name: "make the empty map to slice",
			data: map[int]struct{}{},
			want: []int{},
		},
		{
			name: "make the map to slice",
			data: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, mapKeyToSlice(tc.data))
		})
	}
}

func TestMap(t *testing.T) {
	type user struct {
		id   int
		name string
	}

	testCases := []struct {
		name string
		src  []user
		fn   func(idx int, s user) int
		want []int
	}{
		{
			name: "nil slice",
			src:  nil,
			fn: func(idx int, s user) int {
				return s.id
			},
			want: []int{},
		},
		{
			name: "empty slice",
			src:  []user{},
			fn: func(idx int, s user) int {
				return s.id
			},
			want: []int{},
		},
		{
			name: "non empty slice",
			src: []user{
				{id: 1, name: "Jack"},
				{id: 2, name: "Tom"},
				{id: 3, name: "Gopher"},
			},
			fn: func(idx int, s user) int {
				return s.id
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, Map(tc.src, tc.fn))
		})
	}
}

func TestCombineNestedSlices(t *testing.T) {
	type PostCategory struct {
		Tags []string
	}

	testCases := []struct {
		name        string
		slices      []PostCategory
		extractFunc func(idx int, s PostCategory) []string
		want        []string
	}{
		{
			name: "nil slice",
			want: []string{},
		},
		{
			name:   "empty slice",
			slices: []PostCategory{},
			want:   []string{},
		},
		{
			name: "non empty slice, but nil tags",
			slices: []PostCategory{
				{},
			},
			extractFunc: func(idx int, s PostCategory) []string {
				return s.Tags
			},
			want: []string{},
		},
		{
			name: "non empty slice, but empty tags",
			slices: []PostCategory{
				{Tags: []string{}},
			},
			extractFunc: func(idx int, s PostCategory) []string {
				return s.Tags
			},
			want: []string{},
		},
		{
			name: "non empty slice, but non empty tags",
			slices: []PostCategory{
				{Tags: []string{"Go", "Python", "Kotlin"}},
				{Tags: []string{"Vue", "React", "Golang"}},
			},
			extractFunc: func(idx int, s PostCategory) []string {
				return s.Tags
			},
			want: []string{"Go", "Python", "Kotlin", "Vue", "React", "Golang"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, CombineNestedSlices(tc.slices, tc.extractFunc))
		})
	}
}

func TestCombineAndDeduplicateNestedSlices(t *testing.T) {
	type PostCategory struct {
		Tags []string
	}
	testCases := []struct {
		name        string
		slices      []PostCategory
		extractFunc func(idx int, s PostCategory) []string
		want        []string
	}{
		{
			name: "nil slice",
			want: []string{},
		},
		{
			name:   "empty slice",
			slices: []PostCategory{},
			want:   []string{},
		},
		{
			name: "non empty slice, but nil tags",
			slices: []PostCategory{
				{},
			},
			extractFunc: func(idx int, s PostCategory) []string {
				return s.Tags
			},
			want: []string{},
		},
		{
			name: "non empty slice, but empty tags",
			slices: []PostCategory{
				{Tags: []string{}},
			},
			extractFunc: func(idx int, s PostCategory) []string {
				return s.Tags
			},
			want: []string{},
		},
		{
			name: "non empty slice, but non empty tags",
			slices: []PostCategory{
				{Tags: []string{"Go", "Python", "Kotlin"}},
				{Tags: []string{"Vue", "React", "Golang"}},
			},
			extractFunc: func(idx int, s PostCategory) []string {
				return s.Tags
			},
			want: []string{"Go", "Python", "Kotlin", "Vue", "React", "Golang"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, CombineAndDeduplicateNestedSlices(tc.slices, tc.extractFunc))
		})
	}
}

func TestCombineAndDeduplicateNestedSlicesByEqFunc(t *testing.T) {
	type Tag struct {
		Id   int
		Name string
	}
	type PostCategory struct {
		Tags []Tag
	}
	testCases := []struct {
		name        string
		slices      []PostCategory
		extractFunc func(idx int, s PostCategory) []Tag
		equalFunc   func(a, b Tag) bool
		want        []Tag
	}{
		{
			name: "nil slice",
			want: []Tag{},
		},
		{
			name:   "empty slice",
			slices: []PostCategory{},
			want:   []Tag{},
		},
		{
			name: "non empty slice, but nil tags",
			slices: []PostCategory{
				{},
			},
			extractFunc: func(idx int, s PostCategory) []Tag {
				return s.Tags
			},
			equalFunc: func(a, b Tag) bool {
				return a.Name == b.Name
			},
			want: []Tag{},
		},
		{
			name: "non empty slice, but empty tags",
			slices: []PostCategory{
				{Tags: []Tag{}},
			},
			extractFunc: func(idx int, s PostCategory) []Tag {
				return s.Tags
			},
			equalFunc: func(a, b Tag) bool {
				return a.Name == b.Name
			},
			want: []Tag{},
		},
		{
			name: "non empty slice, but non empty tags",
			slices: []PostCategory{
				{Tags: []Tag{{Id: 1, Name: "Go"}, {Id: 2, Name: "Python"}, {Id: 3, Name: "Kotlin"}}},
				{Tags: []Tag{{Id: 4, Name: "Vue"}, {Id: 5, Name: "React"}, {Id: 6, Name: "Go"}}},
			},
			extractFunc: func(idx int, s PostCategory) []Tag {
				return s.Tags
			},
			equalFunc: func(a, b Tag) bool {
				return a.Name == b.Name
			},
			want: []Tag{
				{Id: 1, Name: "Go"},
				{Id: 2, Name: "Python"},
				{Id: 3, Name: "Kotlin"},
				{Id: 4, Name: "Vue"},
				{Id: 5, Name: "React"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, CombineAndDeduplicateNestedSlicesByEqFunc(tc.slices, tc.extractFunc, tc.equalFunc))
		})
	}
}

func TestIndexStructsByKey(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	testCases := []struct {
		name        string
		slices      []User
		extractFunc func(User) int
		want        map[int]User
	}{
		{
			name: "nil slice",
			want: map[int]User{},
		},
		{
			name:   "empty slice",
			slices: []User{},
			want:   map[int]User{},
		},
		{
			name:   "non empty slice",
			slices: []User{},
			want:   map[int]User{},
		},
		{
			name: "non empty slice",
			slices: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Tom"},
				{Id: 3, Name: "Gopher"},
			},
			extractFunc: func(u User) int {
				return u.Id
			},
			want: map[int]User{
				1: {Id: 1, Name: "John"},
				2: {Id: 2, Name: "Tom"},
				3: {Id: 3, Name: "Gopher"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, IndexStructsByKey(tc.slices, tc.extractFunc))
		})
	}
}

func TestFilterMap(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	testCases := []struct {
		name   string
		slices []User
		fn     func(idx int, s User) (int, bool)
		want   []int
	}{
		{
			name:   "nil slice",
			slices: nil,
			fn: func(idx int, s User) (int, bool) {
				return s.Id, true
			},
			want: []int{},
		},
		{
			name:   "empty slice",
			slices: []User{},
			fn: func(idx int, s User) (int, bool) {
				return s.Id, true
			},
			want: []int{},
		},
		{
			name: "non empty slice",
			slices: []User{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Tom"},
				{Id: 3, Name: "Gopher"},
			},
			fn: func(idx int, s User) (int, bool) {
				if s.Name == "Gopher" {
					return s.Id, true
				}
				return 0, false
			},
			want: []int{3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, FilterMap(tc.slices, tc.fn))
		})
	}
}
