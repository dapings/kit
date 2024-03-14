package std

import (
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
			//rootDomain: "pknic.net.xxx", // old
			rootDomain: "net.xxx", // publicsuffix
		},
		{
			name:   "11.example0.debia.net",
			domain: "11.example0.debian.net",
			//rootDomain: "debian.net", // old
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
			//rootDomain: "y.cn.ye", // old
			rootDomain: "cn.ye", // publicsuffix
		},
		{
			name:   "www.kpu.go.id",
			domain: "www.kpu.go.id",
			//rootDomain: "go.id", // old
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
			//rootDomain: "in.space.museum", // old
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
			//rootDomain: "0emm.com", // old
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
			//rootDomain: "dyndns.org", // old
			rootDomain: "foo.dyndns.org", // publicsuffix
		},
		{
			name:   "foo.blogspot.co.uk",
			domain: "foo.blogspot.co.uk",
			//rootDomain: "blogspot.co.uk", // old
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
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := GetRootDomainName(tt.domain)
			//got := GetRootDomainNameByRegexp(tt.domain)

			if got != tt.rootDomain {
				t.Fatalf("expected %s parts, but got %s", tt.rootDomain, got)
			}
		})
	}
}

func TestSafeGo(t *testing.T) {
	done := make(chan int)
	SafeGo(func() {
		// e.g. http request
		result := 1
		//time.Sleep(3 * time.Second)

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
