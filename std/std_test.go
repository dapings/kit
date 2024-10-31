package std

import (
	"fmt"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestGetWildcardDomain(t *testing.T) {
	testCases := []struct {
		name     string
		domain   string
		wildcard string
	}{
		{
			name:     "empty",
			domain:   "",
			wildcard: "",
		},
		{
			name:     ".cn",
			domain:   ".cn",
			wildcard: "",
		},
		{
			name:     ".com.cn",
			domain:   ".com.cn",
			wildcard: ".com.cn",
		},
		{
			name:     "a.b.com.cn",
			domain:   "a.b.com.cn",
			wildcard: ".b.com.cn",
		},
		{
			name:     "foo.cn",
			domain:   "foo.cn",
			wildcard: "",
		},
		{
			name:     "foo.com.cn",
			domain:   "foo.com.cn",
			wildcard: ".com.cn",
		},
		{
			name:     "a.b.c.com",
			domain:   "a.b.c.com",
			wildcard: ".b.c.com",
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := GetWildcardDomain(tt.domain)

			if got != tt.wildcard {
				t.Fatalf("wildcard domain expected %s, but got %s", tt.wildcard, got)
			}
		})
	}
}

func TestGetRootDomainName(t *testing.T) {
	testCases := []struct {
		name       string
		domain     string
		rootDomain string
	}{
		{
			name:       "empty",
			domain:     "",
			rootDomain: "",
		},
		{
			name:       "www.y.net.ye",
			domain:     "www.y.net.ye",
			rootDomain: "y.net.ye",
		},
		{
			name:       "www.gcom.net.au",
			domain:     "www.gcom.net.au",
			rootDomain: "gcom.net.au",
		},
		{
			name:       "www.nic.net.ge",
			domain:     "www.nic.net.ge",
			rootDomain: "nic.net.ge",
		},
		{
			name:       "dns.marnet.net.mk",
			domain:     "dns.marnet.net.mk",
			rootDomain: "marnet.net.mk",
		},
		{
			name:       "www.monic.net.mo",
			domain:     "www.monic.net.mo",
			rootDomain: "monic.net.mo",
		},
		{
			name:       "cenpac.net.nr",
			domain:     "cenpac.net.nr",
			rootDomain: "cenpac.net.nr",
		},
		{
			name:       "cenpac.net.nz",
			domain:     "cenpac.net.nz",
			rootDomain: "cenpac.net.nz",
		},
		{
			name:       "pk5.pknic.net.pk",
			domain:     "pk5.pknic.net.pk",
			rootDomain: "pknic.net.pk",
		},
		{
			name:       "pk5.pknic.net.sa",
			domain:     "pk5.pknic.net.sa",
			rootDomain: "pknic.net.sa",
		},
		{
			name:       "pk5.pknic.net.sb",
			domain:     "pk5.pknic.net.sb",
			rootDomain: "pknic.net.sb",
		},
		{
			name:       "pk5.pknic.net.sg",
			domain:     "pk5.pknic.net.sg",
			rootDomain: "pknic.net.sg",
		},
		{
			name:   "pk5.pknic.net.xxx",
			domain: "pk5.pknic.net.xxx",
			// rootDomain: "pknic.net.xxx", // old
			rootDomain: "net.xxx", // publicsuffix
		},
		{
			name:   "11.example0.debia.net",
			domain: "11.example0.debian.net",
			// rootDomain: "debian.net", // old
			rootDomain: "example0.debian.net", // publicsuffix
		},
		{
			name:       "example1.debian.org",
			domain:     "example1.debian.org",
			rootDomain: "debian.org",
		},
		{
			name:   "www.y.cn.ye",
			domain: "www.y.cn.ye",
			// rootDomain: "y.cn.ye", // old
			rootDomain: "cn.ye", // publicsuffix
		},
		{
			name:   "www.kpu.go.id",
			domain: "www.kpu.go.id",
			// rootDomain: "go.id", // old
			rootDomain: "kpu.go.id", // publicsuffix
		},
		{
			name:       "www.bar..com.cn",
			domain:     "www.bar.com.cn",
			rootDomain: "bar.com.cn",
		},
		{
			name:       "www.taobao.com",
			domain:     "www.taobao.com",
			rootDomain: "taobao.com",
		},
		{
			name:       "foo.bar.golang.org",
			domain:     "foo.bar.golang.org",
			rootDomain: "golang.org",
		},
		{
			name:       "play.golang.org",
			domain:     "play.golang.org",
			rootDomain: "golang.org",
		},
		{
			name:       "play.golang.net",
			domain:     "play.golang.net",
			rootDomain: "golang.net",
		},
		{
			name:       "play.golang.dev",
			domain:     "play.golang.dev",
			rootDomain: "golang.dev",
		},

		{
			name:   "gophers.in.space.museum",
			domain: "gophers.in.space.museum",
			// rootDomain: "in.space.museum", // old
			rootDomain: "in.space.museum", // publicsuffix
		},
		{
			name:       "amazon.com",
			domain:     "amazon.com",
			rootDomain: "amazon.com",
		},
		{
			name:       "amazon.co.uk",
			domain:     "amazon.co.uk",
			rootDomain: "amazon.co.uk",
		},
		{
			name:       "books.amazon.co.uk",
			domain:     "books.amazon.co.uk",
			rootDomain: "amazon.co.uk",
		},
		{
			name:       "www.books.amazon.co.uk",
			domain:     "www.books.amazon.co.uk",
			rootDomain: "amazon.co.uk",
		},
		{
			name:       "0emm.com",
			domain:     "0emm.com",
			rootDomain: "0emm.com",
		},
		{
			name:       "a.0emm.com",
			domain:     "a.0emm.com",
			rootDomain: "0emm.com",
		},
		{
			name:   "b.c.d.0emm.com",
			domain: "b.c.d.0emm.com",
			// rootDomain: "0emm.com", // old
			rootDomain: "c.d.0emm.com", // publicsuffix
		},
		{
			name:       "there.is.no.such-tld",
			domain:     "there.is.no.such-tld",
			rootDomain: "no.such-tld",
		},
		{
			name:       "foo.org",
			domain:     "foo.org",
			rootDomain: "foo.org",
		},
		{
			name:       "foo.co.uk",
			domain:     "foo.co.uk",
			rootDomain: "foo.co.uk",
		},
		{
			name:   "foo.dyndns.org",
			domain: "foo.dyndns.org",
			// rootDomain: "dyndns.org", // old
			rootDomain: "foo.dyndns.org", // publicsuffix
		},
		{
			name:   "foo.blogspot.co.uk",
			domain: "foo.blogspot.co.uk",
			// rootDomain: "blogspot.co.uk", // old
			rootDomain: "foo.blogspot.co.uk", // publicsuffix
		},
		{
			name:       "cromulent",
			domain:     "cromulent",
			rootDomain: "cromulent",
		},
		{
			name:       "directhr.cn",
			domain:     "directhr.cn",
			rootDomain: "directhr.cn",
		},
		{
			name:       "httpbin.org, used publicsuffix.EffectiveTLDPlusOne err, using regex",
			domain:     "httpbin.org",
			rootDomain: "httpbin.org",
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := GetRootDomainName(tt.domain)
			// got := GetRootDomainNameByRegexp(tt.domain)

			if got != tt.rootDomain {
				t.Fatalf("expected %s parts, but got %s", tt.rootDomain, got)
			}

			t.Logf("root domain: %s", got)
		})
	}
}

func TestParseAndFormatLocalTime(t *testing.T) {
	createTimeStr := "1656644728"
	unixTimestamp, convErr := strconv.Atoi(createTimeStr)
	if convErr == nil {
		createTimeStr = TimeFormat(time.Unix(int64(unixTimestamp), 0))
		createTime, parsedErr := ParseLocalTime(createTimeStr)

		var layoutTimeStr string
		if parsedErr == nil {
			layoutTimeStr = TimeFormat(createTime)
		} else {
			t.Errorf("falied to parse local time, err: %v", parsedErr)
			return
		}
		if layoutTimeStr == "" {
			t.Errorf("expected %s, but got ''", layoutTimeStr)
			return
		}
		t.Logf("lauout time str: %s", layoutTimeStr)
	} else {
		t.Errorf("failed to str conv to int, err: %v", convErr)
	}
}

func TestSafeGo(t *testing.T) {
	done := make(chan int)
	SafeGo(func() {
		// e.g. http request
		result := 1
		// time.Sleep(3 * time.Second)

		select {
		case done <- result:
			println("input result to done chan")
			return
		default:
			println("default ...")
			return
		}
	})

	timeoutTimer := time.NewTimer(5 * time.Second)
	defer timeoutTimer.Stop()

	select {
	case <-done:
		println("output result from done chan")
		return
	case <-timeoutTimer.C:
		println("output timeout from timer chan")
		return
	}
}

func TestFilterPrivateIPs(t *testing.T) {
	ipList := []string{
		"192.168.1.1",  // 私有IPv4
		"10.0.0.1",     // 私有IPv4
		"172.16.254.1", // 私有IPv4
		"2001:0db8:85a3:0000:0000:8a2e:0370:7334", // 公网IPv6
		"fc00::1234", // 私有IPv6
		"fe80::1",    // 链路本地IPv6
		"1.1.1.1",
		"1.1.1.10",
		"123.123.123.123",
		"100.64.0.0",
	}

	expectPrivateNum := 4
	privateIPs := FilterPrivateIPs(ipList)
	if len(privateIPs) != expectPrivateNum {
		t.Errorf("private ips: expect %d, but got %d", expectPrivateNum, len(privateIPs))
		return
	}

	for _, p := range privateIPs {
		t.Log(p)
	}
}

func TestIPsInSubnets(t *testing.T) {
	ips := []string{"192.168.1.10", "192.168.1.11", "10.0.0.1", "2001:db8:a0b:12f0::1", "1.1.1.1", "100.64.0.0"}
	subnets := []string{"192.168.1.0/24", "10.0.0.0/8", "2001:db8:a0b:12f0::1/32", "100.64.0.0/10"}

	inSubnets, err := IPsInSubnets(ips, subnets)
	if err != nil {
		t.Errorf("ips in subnets: %#v", err)
		return
	}

	for ip, inSubnetsResult := range inSubnets {
		t.Logf("ip %s in subnets: %+v", ip, inSubnetsResult)
	}
}

func TestFilterExistingIPsInSubnets(t *testing.T) {
	ips := []string{"192.168.1.10", "192.168.1.11", "10.0.0.1", "2001:db8:a0b:12f0::1", "1.1.1.1", "100.64.0.0"}
	subnets := []string{"192.168.1.0/24", "10.0.0.0/8", "2001:db8:a0b:12f0::1/32", "100.64.0.0/10"}

	invalidAndExistingSourceIPAddr, err := FilterExistingIPsInSubnets(ips, subnets)
	if err != nil {
		t.Errorf("filter existing ips in subnets: %#v", err)
		return
	}

	expectIPAddrNum := 5
	if len(invalidAndExistingSourceIPAddr) != expectIPAddrNum {
		t.Errorf("filter existing ips: expect %d, but got %d", expectIPAddrNum, len(invalidAndExistingSourceIPAddr))
		return
	}

	for _, invalidAndExistingResult := range invalidAndExistingSourceIPAddr {
		t.Logf("invalid and existing ip: %s", invalidAndExistingResult)
	}
}

func TestNetURLParse(t *testing.T) {
	testCases := []string{
		"http://8.137.71.109:8080/auth",
		"https://example.com:8000/auth",
		"https://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:17000/auth",
		"https://[fb11::80]/auth",
		"https://fb11::80/auth",
		"https://[fb11:]:80/auth",
	}

	for i, tt := range testCases {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			val := tt
			addr, err := url.Parse(val)
			if err != nil {
				t.Fatal(err)
			}
			t.Log("host", addr.Host)
			t.Log("hostname", addr.Hostname())
			t.Log("port", addr.Port())

			addr, err = url.ParseRequestURI(val)
			if err != nil {
				t.Fatal(err)
			}
			t.Log("host", addr.Host)
			t.Log("hostname", addr.Hostname())
			t.Log("port", addr.Port())
		})
	}
}

func TestVerifyURLHostIpAddr(t *testing.T) {
	testCases := []string{
		"http://8.137.71.109:8080/auth",
		"https://example.com:8000/auth", // example.com
		"https://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:17000/auth",
		"https://[fb11::80]/auth",  // fb11::80
		"https://[fb11::]:80/auth", // fb11::
		"https://[fb11:]:80/auth",  // invalid
		"https://fb11::80/auth",    // fb11::80
		"https://::1/auth",         // ::1
		"https://::1:80/",          // ::1:80
		"",                         // invalid
	}

	for i, tt := range testCases {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			val := tt
			addr := ValidURLHostIpAddr(val)
			if addr == "" {
				t.Logf("invalid url host ip addr: %q", val)
			} else {
				t.Log(addr)
			}
		})
	}
}
