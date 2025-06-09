package syncx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapKeyLock_Lock(t *testing.T) {
	lock := NewMapKeyLock()
	key1, key2 := "key1", "key2"

	lock.Lock(key1)
	assert.False(t, lock.TryLock(key1))
	assert.False(t, lock.TryRLock(key1))

	assert.True(t, lock.TryLock(key2))

	lock.Unlock(key1)

	assert.True(t, lock.TryLock(key1))
	defer lock.Unlock(key1)

	lock.Unlock(key2)
}

func TestMapKeyLock_RLock(t *testing.T) {
	lock := NewMapKeyLock()
	key1, key2 := "key1", "key2"

	lock.RLock(key1)
	assert.False(t, lock.TryLock(key1))
	assert.True(t, lock.TryRLock(key1))

	lock.RLock(key2)

	lock.RUnlock(key1)

	assert.False(t, lock.TryLock(key1))
	lock.RUnlock(key1)

	assert.True(t, lock.TryLock(key1))
	defer lock.Unlock(key1)

	lock.RUnlock(key2)
}
