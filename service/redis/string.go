package redis

import (
	"fmt"
	"strconv"
)

func (p *Pool) Get(key string) (any, error) {
	result, err := p.Do("GET", key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pool) GetSet(key, value string) (any, error) {
	result, err := p.Do("GETSET", key, value)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pool) Set(key, value string) error {
	_, err := p.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) SetEx(key, value string, seconds int) error {
	_, err := p.Do("SETEX", key, seconds, value)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) SetNx(key, value string) (int64, error) {
	result, err := p.Do("SETNX", key, value)
	if err != nil {
		return 0, err
	}

	if reply, ok := result.(int64); ok {
		return reply, nil
	}

	return 1, nil
}

func (p *Pool) Expire(key string, t int) error {
	_, err := p.Do("EXPIRE", key, t)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) Del(key string) error {
	_, err := p.Do("DEL", key)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) TTL(key string) (int64, error) {
	result, err := p.Do("TTL", key)
	if err != nil {
		return 0, err
	}

	id, ok := result.(int64)
	if !ok {
		err = fmt.Errorf("%s value is not an int64", key)
	}

	return id, err
}

func (p *Pool) Incr(key string) (int64, error) {
	result, err := p.Do("INCR", key)
	if err != nil {
		return 0, err
	}

	id, ok := result.(int64)
	if !ok {
		err = fmt.Errorf("%s value is not an int64", key)
	}

	return id, err
}

func (p *Pool) GetTime() (second, microSecond int64, err error) {
	result, err := p.Do("TIME")
	if err != nil {
		return
	}

	elems, ok := result.([]any)
	if !ok {
		err = EmptyAnswer
		return
	}

	if len(elems) != 2 {
		err = fmt.Errorf("time answer wrong element content %+v", elems)
		return
	}

	secondBytes, ok := elems[0].([]byte)
	if !ok {
		err = fmt.Errorf("type asser into []byte fail %+v", elems[0])
		return
	}
	second, err = strconv.ParseInt(string(secondBytes), 10, 64)
	if err != nil {
		return
	}

	microSecondBytes, ok := elems[1].([]byte)
	if !ok {
		err = fmt.Errorf("type asser into []byte fail %+v", elems[1])
		return
	}

	microSecond, err = strconv.ParseInt(string(microSecondBytes), 10, 64)
	if err != nil {
		return
	}

	return
}
