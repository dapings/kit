package std

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"net"
	"regexp"
	"strings"
)

func Stringify(v ...any) string {
	return fmt.Sprintf(strings.Repeat("%+v", len(v)), v...)
}

func GetMD5(data string) string {
	md5sum := md5.New()
	md5sum.Write([]byte(data))
	return hex.EncodeToString(md5sum.Sum(nil))
}

// ValidIPAddr reports whether a valid ip addr.
func ValidIPAddr(ip string) bool {
	return net.ParseIP(ip) != nil
}

// GetLocalHostIP returns local IP addr.
func GetLocalHostIP() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrList {
		// check the IP addr for loop back addr
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func GetWildcardDomain(domainName string) string {
	if strings.Count(domainName, ".") < 2 {
		// 只有三级或者三级以上的域名才允许有泛域名
		return ""
	}
	if domainName != "" && domainName[0] == '.' {
		return domainName
	}
	return "." + strings.SplitN(domainName, ".", 2)[1]
}

// GetRootDomainName 获取域名的根域名，如果pulicsuffix获取没有出错，返回eTLD，否则使用正则获取
func GetRootDomainName(domainName string) string {
	var result = domainName
	if eTLD, err := publicsuffix.EffectiveTLDPlusOne(strings.TrimPrefix(domainName, ".")); err == nil {
		result = eTLD
	} else {
		// 使用正则兜底
		result = GetRootDomainNameByRegexp(domainName)
	}

	return result
}

// GetRootDomainNameByRegexp 使用正则获取域名的要域名
func GetRootDomainNameByRegexp(domainName string) string {
	var result = domainName
	reg := regexp.MustCompile(`.*\.(com|net|org|gov|edu|cn|co)\.[^.]+$`)
	regRegion := regexp.MustCompile(`.*\.(ac|bj|sh|tj|cq|he|sx|nm|ln|jl|hl|js|zj|ah|fj|jx|sd|ha|hb|hn|gd|gx|hi|sc|gz|yn|gs|qh|nx|xj|sn|xz|mo|hk|tw)\.cn$`)
	if reg.MatchString(domainName) {
		newReg := regexp.MustCompile(`([^.]+\.[^.]+\.[^.]+$)`)
		if newReg.MatchString(domainName) {
			result = newReg.FindString(domainName)
		}
	} else if regRegion.MatchString(domainName) {

		newReg := regexp.MustCompile(`([^.]+\.[^.]+\.cn$)`)
		if newReg.MatchString(domainName) {
			result = newReg.FindString(domainName)
		}
	} else {
		newReg := regexp.MustCompile(`([^.]+\.[^.]+$)`)
		if newReg.MatchString(domainName) {
			result = newReg.FindString(domainName)
		}
	}
	return result
}

// ValidJSON checks whether a str is JSON object, {} or [].
func ValidJSON(str string) bool {
	str = strings.TrimSpace(str)
	if (strings.HasPrefix(str, "{") || strings.HasPrefix(str, "[")) && json.Valid([]byte(str)) {
		return true
	}

	return false
}

// JSONConvert uses json Marshal, Unmarshal to copy.
func JSONConvert(from any, to any) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, to)
}

// JSONDump returns json Marshal string of val.
func JSONDump(val any) string {
	if data, err := json.Marshal(val); err == nil {
		return string(data)
	}

	return ""
}

func SafeGo(run func()) {
	var routine func()

	routine = func() {
		defer func() {
			if err := recover(); err != nil {
				go routine()
			}
		}()

		run()
	}

	go routine()
}
