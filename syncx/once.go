package syncx

import (
	"sync"
	"sync/atomic"
)

// Once is similar to sync.Once, and overrides the Do and doSlow methods with the error return value.
// Once 类似 sync.Once，在其基础上重写 Do 和 doSlow 方法，新增 error 返回值.
type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 0 {
		return o.doSlow(f)
	}
	return nil
}

func (o *Once) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	if atomic.LoadUint32(&o.done) == 0 {
		if err := f(); err != nil {
			return err
		}
		atomic.StoreUint32(&o.done, 1)
	}
	return nil
}
