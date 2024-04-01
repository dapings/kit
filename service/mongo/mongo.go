package mongo

import "strings"

// RewriteKey 由于MongoDB特性限制，对key进行改写
// Field names cannot contain dots or null characters, and must not start with a dollar sign.
func RewriteKey(key string) string {
	if strings.Contains(key, ".") {
		key = strings.Replace(key, ".", "_", -1)
	}

	if len(key) > 0 && key[0] == '$' {
		key = "_" + key[1:]
	}

	return key
}
