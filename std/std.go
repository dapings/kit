package std

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/publicsuffix"
)

var (
	workdir string
)

func Stringify(v ...any) string {
	return fmt.Sprintf(strings.Repeat("%+v", len(v)), v...)
}

func GetMD5(data string) string {
	md5sum := md5.New()
	md5sum.Write([]byte(data))
	return hex.EncodeToString(md5sum.Sum(nil))
}

func GetWorkDir() string {
	if workdir != "" {
		return workdir
	}
	// NOTE: when go run, os.Args[0] is a temp dir, not found files in workdir.
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	dir, _ := filepath.Split(path)
	workdir = strings.Replace(dir, "bin/", "", 1)
	return workdir
}

func WritePidToFile(pidFile string) {
	fd, err := os.OpenFile(pidFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return
	}
	defer fd.Close()

	_, _ = fd.Write([]byte(fmt.Sprintf("%d", os.Getpid())))
}

// MakeNotExists 判断目录是否存在，若不存在，则创建目录
func MakeNotExists(dir string) {
	fd, err := os.Open(dir)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}
	defer fd.Close()
}

func TimeFormat(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}
func ParseLocalTime(timeStr string) (t time.Time, err error) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(timeLayout, timeStr, loc)
}

func ParseLocalDate(dateStr string) (t time.Time, err error) {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(timeLayout, dateStr, loc)
}

// ValidIPAddr reports whether a valid ip addr.
func ValidIPAddr(ip string) bool {
	return net.ParseIP(ip) != nil
}

// ValidURLHostIpAddr 读取URL中的Host IP地址的有效信息
// 如果无效信息或解析错误(url.Parse错误、不是ipv4、不是ipv6、不是有效的域名)，则返回空字符串
func ValidURLHostIpAddr(urlStr string, ipAddrList ...string) string {
	verifyIpAddr := ipAddrList
	if len(ipAddrList) == 0 {
		u, err := url.Parse(urlStr)
		if err != nil {
			return ""
		}
		verifyIpAddr = make([]string, 0, 2)
		host := u.Host
		verifyIpAddr = append(verifyIpAddr, host)
		if host != u.Hostname() {
			// 如果hostname和host不一致，则将hostname加入到verifyIpAddr中
			verifyIpAddr = append(verifyIpAddr, u.Hostname())
		}
	}
	validIpAddr := func(ip string) bool {
		if val := net.ParseIP(ip); val != nil {
			if val.To4() != nil || val.To16() != nil {
				return true
			}
		}
		if IsValidDomainName(ip) {
			return true
		}
		return false
	}
	for _, ip := range verifyIpAddr {
		if validIpAddr(ip) {
			return ip
		}
	}
	return ""
}

// FilterPrivateIPs 过滤出私有IP地址
func FilterPrivateIPs(ips []string) []net.IP {
	var privateIps []net.IP
	for _, ipStr := range ips {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			continue
		}
		if ip.IsLoopback() || ip.IsUnspecified() {
			// 过滤掉回环地址和未指定地址
			continue
		}
		if ip.IsPrivate() {
			privateIps = append(privateIps, ip)
		}
	}
	return privateIps
}

// IPInSubnet 检查单个IP是否在子网内
func IPInSubnet(ip net.IP, subnet *net.IPNet) bool {
	return subnet.Contains(ip)
}

// IPsInSubnets 检查多个IP地址段包含多个IP地址
func IPsInSubnets(ips []string, subnets []string) (map[string][]bool, error) {
	// 预先解析所有子网
	var ipNets []*net.IPNet
	for _, subnetStr := range subnets {
		_, subnet, err := net.ParseCIDR(subnetStr)
		if err != nil {
			return nil, fmt.Errorf("invalid subnet: %s", subnetStr)
		}
		ipNets = append(ipNets, subnet)
	}
	// 并发检查每个IP
	var mu sync.Mutex
	result := make(map[string][]bool)
	var wg sync.WaitGroup
	for _, ipStr := range ips {
		wg.Add(1)
		go func(ipStr string) {
			defer wg.Done()
			ip := net.ParseIP(ipStr)
			mu.Lock()
			if ip == nil {
				// 无效的IP
				result[ipStr] = []bool{}
			}
			result[ipStr] = make([]bool, len(ipNets))
			for i, subnet := range ipNets {
				result[ipStr][i] = IPInSubnet(ip, subnet)
			}
			mu.Unlock()
		}(ipStr)
	}
	wg.Wait()
	return result, nil
}

// FilterExistingIPsInSubnets 过滤存在IP网段中的ip
func FilterExistingIPsInSubnets(ipSources, ipSubnets []string) (result []string, err error) {
	var resultIPsInSubnets map[string][]bool
	resultIPsInSubnets, err = IPsInSubnets(ipSources, ipSubnets)
	if err != nil {
		return
	}
	invalidAndExistingSourceIPAddrInfos := make(map[string]struct{})
	for ip, existingResultInSubnet := range resultIPsInSubnets {
		if _, ok := invalidAndExistingSourceIPAddrInfos[ip]; !ok {
			if len(existingResultInSubnet) == 0 {
				// 无效的ip
				invalidAndExistingSourceIPAddrInfos[ip] = struct{}{}
				continue
			}
			for _, existing := range existingResultInSubnet {
				if existing {
					// ip 存在于某个子网(网段)
					invalidAndExistingSourceIPAddrInfos[ip] = struct{}{}
					break
				}
			}
		}
	}
	var invalidAndExistingSourceIPAddr []string
	for ipAddr := range invalidAndExistingSourceIPAddrInfos {
		invalidAndExistingSourceIPAddr = append(invalidAndExistingSourceIPAddr, ipAddr)
	}
	return invalidAndExistingSourceIPAddr, nil
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
