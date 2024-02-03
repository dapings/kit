package redis

import "github.com/gomodule/redigo/redis"

const (
	ScriptHPop = `
		local v = redis.call("HGET", KEYS[1], ARGV[1])
		redis.call("HDEL", KEYS[1], ARGV[1])
		return v
	`
)

func (p *Pool) HSet(hTable, key, val string) (int64, error) {
	result, err := p.Do("HSET", hTable, key, val)
	if err != nil {
		return 0, err
	}

	return result.(int64), nil
}

func (p *Pool) HSetNX(hTable, key, val string) (int64, error) {
	result, err := p.Do("HSETNX", hTable, key, val)
	if err != nil {
		return 0, err
	}

	return result.(int64), nil
}

func (p *Pool) HGet(hTable, key string) (any, error) {
	result, err := p.Do("HGET", hTable, key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pool) HLen(hTable string) (any, error) {
	result, err := p.Do("HLEN", hTable)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pool) HPop(hTable, key string) (any, error) {
	script := redis.NewScript(1, ScriptHPop)
	result, err := script.Do(p.pool.Get(), hTable, key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pool) HKeys(hTable string) ([]string, error) {
	result, err := p.Do("HKEYS", hTable)
	if err != nil {
		return nil, err
	}

	elems, ok := result.([]any)
	if !ok {
		return nil, WrongAnswer
	}

	results := make([]string, 0, len(elems))
	for _, elem := range elems {
		e, ok := elem.([]byte)
		if !ok {
			continue
		}

		results = append(results, string(e))
	}

	return results, nil
}

func (p *Pool) GetHashValues(hTable string, keys []string) ([]string, error) {
	args := []any{hTable}
	for _, k := range keys {
		args = append(args, k)
	}

	result, err := p.Do("HMGET", args...)
	if err != nil {
		return nil, err
	}

	elems, ok := result.([]any)
	if !ok {
		return nil, WrongAnswer
	}

	results := make([]string, 0, len(elems))
	for _, elem := range elems {
		e, ok := elem.([]byte)
		if !ok {
			continue
		}

		results = append(results, string(e))
	}

	return results, nil
}
