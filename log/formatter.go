package log

import (
	"bytes"
	"fmt"
)

func appendKeyVal(buf *bytes.Buffer, key string, val any) {
	if buf.Len() != 0 {
		buf.WriteByte(',')
	}
	buf.WriteString(key)
	buf.WriteByte('=')
	writeVal(buf, val)
}

func writeVal(buf *bytes.Buffer, val any) {
	strVal, ok := val.(string)
	if !ok {
		strVal = fmt.Sprintf("%+v", val)
	}

	buf.WriteString(strVal)
}
