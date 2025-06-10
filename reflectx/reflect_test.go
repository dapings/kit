package reflectx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPtr(t *testing.T) {
	testCases := []struct {
		name  string
		value any

		want bool
	}{
		{
			name:  "int",
			value: 1,
			want:  false,
		},
		{
			name: "*int",
			value: func() any {
				i := 1
				return &i
			}(),
			want: true,
		},
		{
			name: "**int",
			value: func() any {
				i := 1
				j := &i
				return &j
			}(),
			want: true,
		},
		{
			name:  "nil",
			value: nil,
			want:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, IsPtr(tc.value))
		})
	}
}
