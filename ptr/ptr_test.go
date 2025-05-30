package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	v := 1
	assert.Equal(t, &v, ToPtr(v))
}

func TestGetValueOrDefault(t *testing.T) {
	val := new(int)
	*val = 1
	assert.Equal(t, *val, GetValueOrDefault(val))
	assert.Equal(t, "", GetValueOrDefault[string](nil))
}
