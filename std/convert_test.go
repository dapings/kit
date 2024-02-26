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
		})
	}
}
