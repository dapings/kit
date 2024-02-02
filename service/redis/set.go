package redis

import "github.com/gomodule/redigo/redis"

func (p *Pool) AddToSet(key string, val []string) error {
	_, err := p.Do("SADD", redis.Args{}.Add(key).AddFlat(val)...)
	return err
}

func (p *Pool) RmFromSet(key, val string) error {
	_, err := p.Do("SREM", key, val)
	return err
}

func (p *Pool) GetsFromSet(key string) ([]string, error) {
	reply, err := p.Do("SMEMBERS", key)
	if err != nil {
		return nil, err
	}

	elems, ok := reply.([]any)
	if !ok {
		return nil, WrongAnswer
	}

	results := make([]string, 0, len(elems))
	for _, elem := range elems {
		val, ok := elem.([]byte)
		if !ok {
			continue
		}

		results = append(results, string(val))
	}

	return results, nil
}
