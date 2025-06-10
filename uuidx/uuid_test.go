package uuidx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID4(t *testing.T) {
	t.Run("uuid4", func(t *testing.T) {
		uuid := UUID4()
		assert.NotEmpty(t, uuid)
	})
}

func TestRearrangeStrUUID4(t *testing.T) {
	testCases := []struct {
		name    string
		uuid    string
		want    string
		wantErr error
	}{
		{
			name:    "invalid uuid",
			uuid:    "58e0a7d7-eebc-11d8-9669-0800200c9a66-1234",
			want:    "",
			wantErr: errors.New("invalid uuid"),
		},
		{
			name:    "valid uuid",
			uuid:    "58e0a7d7-eebc-11d8-9669-0800200c9a66",
			want:    "11d8eebc58e0a7d796690800200c9a66",
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := RearrangeStrUUID4(tc.uuid)
			if tc.wantErr != nil {
				assert.Equal(t, err, tc.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestRearrangeUUID4(t *testing.T) {
	t.Run("rearrange uuid4", func(t *testing.T) {
		uuid := RearrangeUUID4()
		assert.NotEmpty(t, uuid)
	})
}
