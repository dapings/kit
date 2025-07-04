package dns

import (
	"net"
	"sync"
	"time"
)

var (
	dnsCache      = make(map[string]*cachedIP)
	dnsCacheMutex sync.RWMutex
	cacheTTL      = 1 * time.Minute
)

type cachedIP struct {
	IPs    []string
	Expire time.Time
}

// ResolveAndCache resolves the IP addresses of a host and caches the result.
func ResolveAndCache(host string) ([]string, error) {
	dnsCacheMutex.RLock()
	if cached, ok := dnsCache[host]; ok && cached.Expire.After(time.Now()) {
		dnsCacheMutex.RUnlock()
		return cached.IPs, nil
	}
	dnsCacheMutex.RUnlock()

	dnsCacheMutex.Lock()
	defer dnsCacheMutex.Unlock()

	// 再次检查缓存(双重检查)
	if cached, ok := dnsCache[host]; ok && cached.Expire.After(time.Now()) {
		return cached.IPs, nil
	}

	ips, err := net.LookupHost(host)
	if err != nil {
		return nil, err
	}

	expire := time.Now().Add(cacheTTL)
	dnsCache[host] = &cachedIP{
		IPs:    ips,
		Expire: expire,
	}
	return ips, nil
}
