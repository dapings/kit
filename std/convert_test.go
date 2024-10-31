package std

import (
	"testing"
)

func TestSplitDomainByLevel(t *testing.T) {
	testDomains := []struct {
		name   string
		want   int
		level  int
		domain string
	}{
		{
			name:   "contain tld",
			want:   4,
			level:  3,
			domain: "1.com.1.boinlive1.com",
		},
		{
			name:   "contain tld, default level 1",
			want:   2,
			level:  0,
			domain: "1.com.1.boinlive1.com",
		},
		{
			name:   "contain tld, default level 1",
			want:   2,
			level:  5,
			domain: "1.com.1.boinlive1.com",
		},
		{
			name:   "only contain tld, wildcard tld(level - 1 = 2)",
			want:   2,
			level:  3,
			domain: "1.httpbin.org",
		},
	}

	for _, tt := range testDomains {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			parts, err := SplitDomainBySpecifyLevel(tt.domain, tt.level)
			if err != nil {
				t.Fatal(err)
			}

			if len(parts) != tt.want {
				t.Fatalf("expected %d parts, but got %d", tt.want, len(parts))
			}

			t.Logf("got parts: %v", parts)
		})
	}
}

func TestIsValidDomainName(t *testing.T) {
	ok := IsValidDomainName("www.tmall.com")
	if ok == false {
		t.Fatal("IsValidDomainName")
	}
	ok = IsValidDomainName("wwwtmallcom")
	if ok {
		t.Fatal("IsValidDomainName")
	}
	ok = IsValidDomainName("1.1.1.1")
	if ok {
		t.Fatal("should be invalid domain")
	}
	ok = IsValidDomainName("test.zeus.58")
	if ok {
		t.Fatal("should be invalid domain")
	}
}

func TestRemoveStringSlice(t *testing.T) {
	source := []string{"7", "8", "9", "8", "9", "10", "6", "6", "6", "6"}
	removed := []string{"7", "8", "9"}
	result := RemoveStringSlice(source, removed)
	t.Log(result)

	result = RemoveStringSlice([]string{"10"}, result)
	t.Log(result)

	// 移除重复的
	result = RemoveStringSlice(RemoveDuplicateStringList(source), RemoveDuplicateStringList(removed))
	t.Log(result)

	result = RemoveStringSlice([]string{"10"}, result)
	t.Log(result)
}
