package redis

import (
	"fmt"
	"time"
)

const (
	hTable           = "_htable"
	retryQueueSuffix = "_retry"
	lockKeySuffix    = "_lock"
)

func EnQueueReliably(pool *Pool, queue string, key string, data []byte) (err error) {
	for {
		err = EnQueue(pool, queue, key, data)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return
}

func EnQueue(pool *Pool, queueName, key string, data []byte) error {
	if pool == nil {
		return fmt.Errorf("redis pool is empty")
	}

	if key == "" {
		return fmt.Errorf("queue %s using invalid key: %s", queueName, key)
	}

	result, err := pool.HSet(queueName+hTable, key, string(data))
	if err != nil {
		return err
	}

	// 数据合并，队列不重复添加
	if result == 0 {
		return nil
	}

	err = pool.RPush(queueName, []byte(key))
	if err != nil {
		return err
	}

	return nil
}

func DeQueue(pool *Pool, queueName string) (key string, data []byte) {
	if pool == nil {
		return "", nil
	}

	result, err := pool.LPop(queueName)
	if err != nil || result == nil {
		return "", nil
	}

	k, ok := result.([]byte)
	if !ok {
		return "", nil
	}

	key = string(k)
	result, err = pool.HPop(queueName+hTable, key)
	if err != nil || result == nil {
		return "", nil
	}

	data, ok = result.([]byte)
	if !ok {
		return "", nil
	}

	return
}

func CheckLocked(pool *Pool, key string) bool {
	val, err := pool.Get(key)
	if err != nil || val != nil {
		return true
	}

	return false
}
