package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func (p *Pool) LPush(args ...any) error {
	_, err := p.Do("LPUSH", args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) LRange(key string, start, end int) (any, error) {
	result, err := redis.Values(p.Do("LRANGE", key, start, end))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pool) RPush(args ...any) error {
	_, err := p.Do("RPUSH", args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) LPop(key string) (any, error) {
	result, err := p.Do("LPOP", key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pool) LLen(key string) (int64, error) {
	result, err := p.Do("LLEN", key)
	if err != nil {
		return 0, err
	}

	length, ok := result.(int64)
	if !ok {
		err = fmt.Errorf("%s value is not a int64", key)
	}

	return length, err
}

func (p *Pool) RPop(key string) (any, error) {
	result, err := p.Do("RPOP", key)
	if err != nil {
		return nil, err
	}

	return result, nil
}
