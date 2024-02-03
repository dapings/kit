package redis

import (
	"fmt"
	"time"
)

func Lock(globalPool *Pool, key string, maxRetry, expiredSecond int) error {
	if globalPool == nil {
		return fmt.Errorf("global redis pool is empty")
	}

	// 先竞争锁，抢锁成功后才处理其他
	if maxRetry <= 0 {
		maxRetry = 5
	}
	var i = 0
	for ; i < maxRetry; i++ {
		result, err := globalPool.SetNx(key, "1")
		if err != nil {
			return err
		}
		if result != 1 {
			if maxRetry > 2 {
				time.Sleep(100 * time.Millisecond)
			}

			continue
		}

		break
	}

	if i == maxRetry {
		return fmt.Errorf("lock failed %s", key)
	}

	// 锁不主动释放，设置超时时间后自动释放
	// 设置自动超时，过期回收，避免分布式锁死锁
	if expiredSecond <= 0 {
		expiredSecond = 3600
	}
	var expireHandle = func() error {
		return globalPool.Expire(key, expiredSecond)
	}
	err := expireHandle()
	if err != nil {
		_ = expireHandle()
	}

	return nil
}

// LockByExpireTime 获取锁，可以自定义重试时长、锁超时时间。
func LockByExpireTime(globalPool *Pool, key string, timeout time.Duration, expired int) error {
	if globalPool == nil {
		return fmt.Errorf("global redis pool is empty")
	}

	start := time.Now()
	for {
		now := time.Now()
		if now.Sub(start) > timeout {
			return fmt.Errorf("get %s lock already expire", key)
		}

		result, err := globalPool.SetNx(key, "1")
		if err != nil {
			return err
		}
		if result != 1 {
			continue
		}

		break
	}

	// 设置自动超时，过期回收，避免分布式锁死锁
	var expireHandle = func() error {
		return globalPool.Expire(key, expired)
	}
	err := expireHandle()
	if err != nil {
		_ = expireHandle()
	}

	return nil
}

func UnLock(globalPool *Pool, key string) error {
	if globalPool == nil {
		return fmt.Errorf("global redis pool is empty")
	}

	if key == "" {
		return fmt.Errorf("unlocked key is empty")
	}

	if err := globalPool.Del(key); err != nil {
		return err
	}

	return nil
}
