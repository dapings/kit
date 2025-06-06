package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	testCases := []struct {
		name string
		mp   map[int]int
		want []int
	}{
		{
			name: "nil map",
			mp:   nil,
			want: make([]int, 0),
		},
		{
			name: "empty map",
			mp:   make(map[int]int),
			want: make([]int, 0),
		},
		{
			name: "normal map with 3 element",
			mp:   map[int]int{1: 1, 2: 2, 3: 3},
			want: []int{1, 2, 3},
		},
		{
			name: "normal map with 4 element",
			mp:   map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Keys(tc.mp)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestValues(t *testing.T) {
	testCases := []struct {
		name string
		mp   map[int]string
		want []string
	}{
		{
			name: "nil map",
			mp:   nil,
			want: make([]string, 0),
		},
		{
			name: "empty map",
			mp:   make(map[int]string),
			want: make([]string, 0),
		},
		{
			name: "normal map with 3 element",
			mp:   map[int]string{1: "a", 2: "b", 3: "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "normal map with 4 element",
			mp:   map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Values(tc.mp)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestKeyValues(t *testing.T) {
	testCases := []struct {
		name       string
		mp         map[int]string
		wantKeys   []int
		wantValues []string
	}{
		{
			name:       "nil map",
			mp:         nil,
			wantKeys:   make([]int, 0),
			wantValues: make([]string, 0),
		},
		{
			name:       "empty map",
			mp:         make(map[int]string),
			wantKeys:   make([]int, 0),
			wantValues: make([]string, 0),
		},
		{
			name:       "normal map with 3 element",
			mp:         map[int]string{1: "a", 2: "b", 3: "c"},
			wantKeys:   []int{1, 2, 3},
			wantValues: []string{"a", "b", "c"},
		},
		{
			name:       "normal map with 4 element",
			mp:         map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
			wantKeys:   []int{1, 2, 3, 4},
			wantValues: []string{"a", "b", "c", "d"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotKeys, gotValues := KeyValues(tc.mp)
			assert.ElementsMatch(t, tc.wantKeys, gotKeys)
			assert.ElementsMatch(t, tc.wantValues, gotValues)
		})
	}
}
