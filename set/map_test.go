package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSet_Add(t *testing.T) {
	testCases := []struct {
		name string
		m    MapSet[int]
		val  int

		expected MapSet[int]
	}{
		{
			name: "empty MapSet",
			m:    NewMapSet[int](0),
			val:  1,
			expected: MapSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
		},
		{
			name: "MapSet with the existed elements",
			m: MapSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
			val: 2,
			expected: MapSet[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Add(tt.val)
			assert.Equal(t, tt.expected, tt.m)
		})
	}
}

func TestMapSet_Remove(t *testing.T) {
	testCases := []struct {
		name string
		m    MapSet[int]
		val  int

		expected MapSet[int]
	}{
		{
			name: "empty MapSet",
			m:    NewMapSet[int](0),
			val:  1,
			expected: MapSet[int]{
				m: map[int]struct{}{},
			},
		},
		{
			name: "MapSet with the existed elements",
			m: MapSet[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
				},
			},
			val: 2,
			expected: MapSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Remove(tt.val)
			assert.Equal(t, tt.expected, tt.m)
		})
	}
}

func TestMapSet_Contains(t *testing.T) {
	testCases := []struct {
		name string
		m    MapSet[int]
		val  int

		expected bool
	}{
		{
			name:     "empty MapSet",
			m:        NewMapSet[int](0),
			val:      1,
			expected: false,
		},
		{
			name: "MapSet with the existed elements, not contain",
			m: MapSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
			val:      2,
			expected: false,
		},
		{
			name: "MapSet with the existed elements, contain",
			m: MapSet[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
				},
			},
			val:      2,
			expected: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.m.Contains(tt.val))
		})
	}
}

func TestMapSet_Clear(t *testing.T) {
	testCases := []struct {
		name string
		m    MapSet[int]

		expected MapSet[int]
	}{
		{
			name: "empty MapSet",
			m:    NewMapSet[int](0),
			expected: MapSet[int]{
				m: map[int]struct{}{},
			},
		},
		{
			name: "MapSet with the existed elements",
			m: MapSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
			expected: MapSet[int]{
				m: map[int]struct{}{},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Clear()
			assert.Equal(t, tt.expected, tt.m)
		})
	}
}

func TestMapSet_IsEmpty(t *testing.T) {
	testCases := []struct {
		name string
		m    MapSet[int]

		expected bool
	}{
		{
			name:     "empty MapSet",
			m:        NewMapSet[int](0),
			expected: true,
		},
		{
			name: "MapSet with the existed elements",
			m: MapSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
			expected: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.m.IsEmpty())
		})
	}
}

func TestMapSet_Size(t *testing.T) {
	testCases := []struct {
		name string
		m    MapSet[int]

		expected int
	}{
		{
			name:     "empty MapSet",
			m:        NewMapSet[int](0),
			expected: 0,
		},
		{
			name: "MapSet with the existed elements",
			m: MapSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
			expected: 1,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.m.Size())
		})
	}
}

func TestNewMapSet(t *testing.T) {
	testCases := []struct {
		name string
		cap  int

		expected MapSet[int]
	}{
		{
			name: "MapSet with the initial cap 0",
			cap:  0,
			expected: MapSet[int]{
				m: map[int]struct{}{},
			},
		},
		{
			name: "MapSet with the initial cap 1",
			cap:  1,
			expected: MapSet[int]{
				m: make(map[int]struct{}, 1),
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, NewMapSet[int](tt.cap))
		})
	}
}
