package redis

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	MaxFail      = 6
	FailureRetry = 3

	ServerEmptyErr        = errors.New("redis: servers is empty")
	ServerNonAvailableErr = errors.New("redis: server non available ")

	WrongAnswer = errors.New("redis: get wrong answer")
	EmptyAnswer = errors.New("redis: get empty answer")
)

type Server struct {
	Host string
	Port string
	Auth string
	DB   string
}

type Pool struct {
	mu *sync.RWMutex

	// the redis instance index of pools
	index  int
	pools  []*redis.Pool
	pool   *redis.Pool
	status map[int]bool
	// 熵值策略可以防止抖动引起的主备切换，只有连接失败才进行主备切换
	// 采用熵值[0,MaxFail]记录redis稳定性
	// 连接失败熵值加一加到MaxFail进行切换，连接成功熵值减一减到0为止
	entropy map[int]int
	servers map[int]string
}

func (p *Pool) checkDoTest() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, v := range p.status {
		if v {
			return false
		}
	}

	return true
}

func (p *Pool) testAll() {
	// 当所有的redis都不可用时，尝试重新探测一个可用的
	if p.checkDoTest() {
		for i := range p.pools {
			go func(index int) {
				conn := p.pools[index].Get()
				_, err := conn.Do("PING")
				var opError *net.OpError
				if !errors.As(err, &opError) && p.checkDoTest() {
					p.mu.Lock()
					// switch redis instance to index
					p.switchIndex(index)
					p.mu.Unlock()
				}
			}(i)
		}
	}
}

// switchIndex switch redis instance to index.
func (p *Pool) switchIndex(index int) {
	// NOTE: 调用前记得加锁
	if index >= len(p.pools) {
		return
	}

	// redis switch to index p.servers[index]
	p.index = index
	p.entropy[index] = 0
	p.status[index] = true
	p.pool = p.pools[index]
}

func (p *Pool) FailOver() {
	// 当只有一个server连接时，不进行切换
	if len(p.pools) == 1 {
		return
	}

	// 记录失败连接数，当大于最大失败后切换server
	p.mu.Lock()
	defer p.mu.Unlock()

	p.entropy[p.index] = p.entropy[p.index] + 1
	if p.entropy[p.index] >= MaxFail {
		p.status[p.index] = false
		index := p.index + 1
		if index >= len(p.pools) {
			index = 0
		}
		p.switchIndex(index)
	}

	return
}

func (p *Pool) Recover() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.entropy[p.index] > 0 {
		p.entropy[p.index] = p.entropy[p.index] - 1
	}

	return
}

func (p *Pool) Do(cmdStr string, args ...any) (reply any, err error) {
	// actually do the redis commands
	// 所有server状态异常后，每次都尝试重新测试
	if p.checkDoTest() {
		go p.testAll()
		return nil, ServerNonAvailableErr
	}

	// 失败重试 FailureRetry 次
	var conn redis.Conn
	defer conn.Close()
	for i := 0; i < FailureRetry; i++ {
		conn = p.pool.Get()

		if reply, err = conn.Do(cmdStr, args...); err != nil {
			var opErr *net.OpError
			if errors.As(err, &opErr) {
				p.FailOver()
				go p.testAll()
			}

			time.Sleep(100 * time.Millisecond)

			continue
		}

		// 重试 i 次后成功
		p.Recover()
		break
	}

	return
}

func (p *Pool) Close() error {
	for i := range p.pools {
		p.pools[i].Close()
	}
	p.pool = nil

	return nil
}

func NewRedisPool(s []Server) (*Pool, error) {
	if len(s) <= 0 {
		return nil, ServerEmptyErr
	}

	p := &Pool{
		status:  make(map[int]bool),
		entropy: make(map[int]int),
		servers: make(map[int]string),
	}

	for i := range s {
		rp, _ := newTimeoutPool(s[i], 1*time.Second, 2*time.Second, 2*time.Second)
		if _, err := rp.Dial(); err != nil {
			continue
		}

		p.pools = append(p.pools, rp)
		p.entropy[i] = 0
		p.status[i] = true
		p.servers[i] = s[i].Host
	}

	p.index = 0
	if len(p.pools) > p.index {
		p.pool = p.pools[p.index]
	}

	return p, nil
}

func newTimeoutPool(s Server, connTimeout, readTimeout, writeTimeout time.Duration) (*redis.Pool, error) {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port),
				redis.DialConnectTimeout(connTimeout),
				redis.DialReadTimeout(readTimeout),
				redis.DialWriteTimeout(writeTimeout))
			if err != nil {
				return nil, err
			}

			if !(s.Auth == "" || s.Auth == "nil") {
				_, err = conn.Do("AUTH", s.Auth)
				if err != nil {
					conn.Close()
					return nil, err
				}
			}

			if s.DB != "" {
				_, err = conn.Do("SELECT", s.DB)
				if err != nil {
					conn.Close()
					return nil, err
				}
			}

			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, _ time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
		MaxIdle:     500,
		IdleTimeout: 600 * time.Second,
	}, nil
}
