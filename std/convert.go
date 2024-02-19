package std

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
)

func GetMD5(data string) string {
	md5sum := md5.New()
	md5sum.Write([]byte(data))
	return hex.EncodeToString(md5sum.Sum(nil))
}

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
