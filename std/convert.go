package std

import (
	"regexp"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func RemoveDuplicateStringSlice(tmp []string) []string {
	seen := make(map[string]struct{})
	var result []string
	for _, t := range tmp {
		if _, ok := seen[t]; !ok {
			seen[t] = struct{}{}
			result = append(result, t)
		}
	}

	return result
}

func UnionStringSlice(s1, s2 []string) []string {
	for _, s := range s2 {
		if i := FindStrInSlice(s1, s); i == -1 {
			s1 = append(s1, s)
		}
	}

	return s1
}

func RemoveStringSlice(s []string, removed []string) []string {
	for _, v := range removed {
		if i := FindStrInSlice(s, v); i != -1 {
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}

func FindStrInSlice(list []string, str string) int {
	for i, v := range list {
		if v == str {
			return i
		}
	}

	return -1
}

func ReverseStr(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func LowerFirstChar(letter string) string {
	if len(letter) == 0 {
		return letter
	}

	return strings.ToLower(letter[:1] + letter[1:])
}

func UpperFirstChar(letter string) string {
	if len(letter) == 0 {
		return letter
	}

	return strings.ToUpper(letter[:1] + letter[1:])
}

func IsLetterOrNumber(str string) bool {
	for _, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}

	return true
}

func IsWildcardDomain(domainName string) bool {
	//	泛域名
	return domainName != "" && domainName[0] == '.'
}

func IsValidDomainName(domainName string) bool {
	if IsWildcardDomain(domainName) {
		domainName = domainName[1:]
	}

	domainRegexp := regexp.MustCompile("^([a-z0-9]([a-z0-9\\-]{0,61}[a-z0-9])?\\.)+[a-z][a-z0-9\\-]{0,62}$")
	matches := domainRegexp.FindStringSubmatch(domainName)
	if len(matches) < 1 {
		return false
	}

	return true
}

// SplitDomainBySpecifyLevel 将域名按指定层级拆分，形成泛域名+域名集合
func SplitDomainBySpecifyLevel(domainName string, level int) ([]string, error) {
	levelDomains := []string{domainName}
	if !IsValidDomainName(domainName) {
		return levelDomains, nil
	}

	// 生成 TLD 泛域名
	tld, err := publicsuffix.EffectiveTLDPlusOne(strings.TrimPrefix(domainName, "."))
	if err != nil {
		return levelDomains, err
	}

	levelDomains = append(levelDomains, "."+tld)
	notTLDPartDomain := strings.TrimSuffix(domainName, "."+tld)
	parts := strings.Split(notTLDPartDomain, ".")

	if level <= 0 {
		level = 1
	}
	// remove tld level
	level = level - 1
	if len(parts) <= level-1 {
		return levelDomains, nil
	}

	for i := 0; i < level; i++ {
		levelDomains = append(levelDomains, "."+strings.Join(parts[len(parts)-i-1:], ".")+levelDomains[1])
	}

	// remove dup from levelDomains
	unique := make([]string, 0, len(levelDomains))
	seen := make(map[string]struct{})
	for _, domain := range levelDomains {
		if _, ok := seen[domain]; !ok {
			unique = append(unique, domain)
			seen[domain] = struct{}{}
		}
	}
	levelDomains = unique

	return levelDomains, nil
}
