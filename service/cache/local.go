package cache

import (
	"github.com/dapings/kit/std"
	"golang.org/x/sync/syncmap"
	"time"
)

func Get(k any) (any, bool) {
	item, ok := getItem(k)
	if ok && item != nil {
		return item.Value, ok
	}

	return nil, ok
}

func Set(k, v any, expire time.Duration) {
	item := &CachedItem{
		Expire:     expire,
		CachedTime: time.Now().Unix(),
		Value:      v,
	}

	defaultMemoryCache.m.Store(k, item)
}

func Refresh() {
	f := func(k, v any) bool {
		item, ok := getItem(k)
		if ok && item.Expire > 0 && (time.Now().Unix()-item.CachedTime) > int64(item.Expire/time.Second) {
			defaultMemoryCache.m.Delete(k)
		}

		return true
	}

	defaultMemoryCache.m.Range(f)
}

type CachedItem struct {
	Expire     time.Duration
	CachedTime int64 // a timestamp
	Value      any
}

type Memory struct {
	m syncmap.Map
}

func getItem(k any) (*CachedItem, bool) {
	v, ok := defaultMemoryCache.m.Load(k)
	if ok {
		item, ok := v.(*CachedItem)
		return item, ok
	}

	return nil, ok
}

var (
	defaultMemoryCache *Memory
)

func init() {
	defaultMemoryCache = &Memory{}
	std.SafeGo(func() {
		for {
			Refresh()

			time.Sleep(1 * time.Second)
		}
	})
}
