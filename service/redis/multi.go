package redis

import (
	"errors"
	"net"
)

func (p *Pool) Multi(cmds []map[string]string) (any, error) {
	// 所有 server 实例异常后，每次都尝试重新测试
	if p.checkDoTest() {
		go p.testAll()

		return nil, ServerNonAvailableErr
	}

	conn := p.pool.Get()
	defer conn.Close()

	for _, cmd := range cmds {
		err := conn.Send("WATCH", cmd["key"])
		if err != nil {
			return nil, err
		}
	}

	err := conn.Send("MULTI")
	if err != nil {
		return nil, err
	}
	for _, cmd := range cmds {
		err = conn.Send(cmd["cmdName"], cmd["key"], cmd["value"])
		if err != nil {
			err = conn.Send("DISCARD")
			return nil, err
		}
	}

	var result any
	result, err = conn.Do("EXEC")
	if err != nil {
		var opErr *net.OpError
		if errors.As(err, &opErr) {
			p.FailOver()
			go p.testAll()
		}

		err = conn.Send("DISCARD")
		return nil, err
	}

	p.Recover()

	return result, nil
}

func (p *Pool) MultiVariable(cmds []map[string][]interface{}) (any, error) {
	// 所有 server 实例异常后，每次都尝试重新测试
	if p.checkDoTest() {
		go p.testAll()

		return nil, ServerNonAvailableErr
	}

	conn := p.pool.Get()
	defer conn.Close()

	for _, cmd := range cmds {
		for _, args := range cmd {
			if len(args) < 1 {
				return nil, errors.New("param error")
			}

			err := conn.Send("WATCH", args[0])
			if err != nil {
				return nil, err
			}
		}
	}

	err := conn.Send("MULTI")
	if err != nil {
		return nil, err
	}
	for _, cmd := range cmds {
		for cmdName, args := range cmd {
			err = conn.Send(cmdName, args...)
			if err != nil {
				err = conn.Send("DISCARD")
				return nil, err
			}
		}
	}

	var result any
	result, err = conn.Do("EXEC")
	if err != nil {
		var opErr *net.OpError
		if errors.As(err, &opErr) {
			p.FailOver()
			go p.testAll()
		}

		err = conn.Send("DISCARD")
		return nil, err
	}

	p.Recover()

	return result, nil
}
