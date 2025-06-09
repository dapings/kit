package syncx

import "sync"

type KeyLocker interface {
	Lock(key string)
	Unlock(key string)
}

type MapKeyLock struct {
	locks sync.Map
}

func NewMapKeyLock() *MapKeyLock {
	return &MapKeyLock{}
}

func (l *MapKeyLock) Lock(key string) {
	mu, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	mu.(*sync.RWMutex).Lock()
}

func (l *MapKeyLock) Unlock(key string) {
	if mu, ok := l.locks.Load(key); ok {
		mu.(*sync.RWMutex).Unlock()
	}
}

func (l *MapKeyLock) RLock(key string) {
	mu, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	mu.(*sync.RWMutex).RLock()
}

func (l *MapKeyLock) RUnlock(key string) {
	if mu, ok := l.locks.Load(key); ok {
		mu.(*sync.RWMutex).RUnlock()
	}
}

func (l *MapKeyLock) TryLock(key string) bool {
	mu, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	return mu.(*sync.RWMutex).TryLock()
}

func (l *MapKeyLock) TryRLock(key string) bool {
	mu, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	return mu.(*sync.RWMutex).TryRLock()
}
