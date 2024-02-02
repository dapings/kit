package std

import (
	"encoding/json"
	"net"
	"strings"
)

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
