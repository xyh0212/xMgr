package x

import (
	// "strings"
	"net/url"
)

// reserveBuffer checks cap(buf) and expand buffer to len(buf) + appendSize.
// If cap(buf) is not enough, reallocate new buffer.
func reserveBuffer(buf []byte, appendSize int) []byte {
	newSize := len(buf) + appendSize
	if cap(buf) < newSize {
		// Grow buffer exponentially
		newBuf := make([]byte, len(buf)*2+appendSize)
		copy(newBuf, buf)
		buf = newBuf
	}
	return buf[:newSize]
}

// func EscapeString(str string) string {
// 	str = strings.Replace(str, "\"", "\\\"", -1)
// 	return str
// }
func EscapeString(buf []byte, v string) []byte {
	pos := len(buf)
	buf = reserveBuffer(buf, len(v)*2)

	for i := 0; i < len(v); i++ {
		c := v[i]
		switch c {
		// case '\x00':
		// 	buf[pos] = '\\'
		// 	buf[pos+1] = '0'
		// 	pos += 2
		// case '\n':
		// 	buf[pos] = '\\'
		// 	buf[pos+1] = 'n'
		// 	pos += 2
		// case '\r':
		// 	buf[pos] = '\\'
		// 	buf[pos+1] = 'r'
		// 	pos += 2
		// case '\x1a':
		// 	buf[pos] = '\\'
		// 	buf[pos+1] = 'Z'
		// 	pos += 2
		// case '\'':
		// 	buf[pos] = '\\'
		// 	buf[pos+1] = '\''
		// 	pos += 2
		case '"':
			buf[pos] = '\\'
			buf[pos+1] = '"'
			pos += 2
		// case '\\':
		// 	buf[pos] = '\\'
		// 	buf[pos+1] = '\\'
		// 	pos += 2
		default:
			buf[pos] = c
			pos++
		}
	}

	return buf[:pos]
}

func UrlEncode(buf []byte, v string) []byte {
	// l, _ := url.ParseQuery(v)
	// v = l.Encode()
	v = url.QueryEscape(v)

	pos := len(buf)
	buf = reserveBuffer(buf, len(v)*2)

	for i := 0; i < len(v); i++ {
		c := v[i]
		switch c {
		case '&': //%26
			buf[pos] = '%'
			buf[pos+1] = '2'
			buf[pos+2] = '6'
			pos += 3
		case '=': //%3d
			buf[pos] = '%'
			buf[pos+1] = '3'
			buf[pos+2] = 'd'
			pos += 3
		default:
			buf[pos] = c
			pos++
		}
	}

	return buf[:pos]
}
