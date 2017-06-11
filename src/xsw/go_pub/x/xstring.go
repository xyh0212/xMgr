package x

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func AtoI64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return i, err
}
func SubStr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	if length < 0 {
		length = rl
	}

	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// snake string, XxYy to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// [["123478","aaaa","123456","0",""],
// ["123482","123","12223","0",""]]
func Any2JsonArrayString(i interface{}) (str string) {
	o := any2JstonString{}
	return o.Any2String(i)
}

/////////////////////////////////////////////////////////////
//Any2String
type any2JstonString struct {
	pc        int
	bShowName bool
}

func (self *any2JstonString) getStructString(t reflect.Type, v reflect.Value) (str string) {
	n := t.NumField()
	str += "["

	for i := 0; i < n; i++ {
		value := v.Field(i)
		if self.bShowName {
			str += fmt.Sprint(strings.Repeat(" ", self.pc), t.Field(i).Name, ":")
			str += fmt.Sprintln(strings.Repeat(" ", self.pc), value.Interface())
		} else {
			if i > 0 {
				str += fmt.Sprint(",\"", value.Interface(), "\"")
			} else {
				str += fmt.Sprint("\"", value.Interface(), "\"")
			}
		}
		// Any2String(value.Interface(), pc+2)
	}
	str += "]"
	return str
}

func (self *any2JstonString) getArraySliceString(v reflect.Value) (str string) {
	str += "["
	n := v.Len()
	for j := 0; j < n; j++ {
		str += self.Any2String(v.Index(j).Interface())
		if j < (n - 1) {
			str += ","
		}
	}
	str += "]"
	return str
}

func (self *any2JstonString) getMapString(v reflect.Value) (str string) {
	n := len(v.MapKeys())
	i := 0
	for _, k := range v.MapKeys() {
		// json数组不显示key
		//str += self.Any2String(k.Interface())
		str += self.Any2String(v.MapIndex(k).Interface())
		i += 1
		if i < n {
			str += ","
		}
	}
	return str
}

func (self *any2JstonString) Any2String(i interface{}) (str string) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = reflect.ValueOf(i).Elem()
		t = v.Type()
	}
	switch v.Kind() {
	case reflect.Array:
		str += self.getArraySliceString(v)
	case reflect.Chan:
		str += fmt.Sprintln("Chan")
	case reflect.Func:
		str += fmt.Sprintln("Func")
	case reflect.Interface:
		str += fmt.Sprintln("Interface")
	case reflect.Map:
		str += self.getMapString(v)
	case reflect.Slice:
		str += self.getArraySliceString(v)
	case reflect.Struct:
		str += self.getStructString(t, v)
	case reflect.UnsafePointer:
		str += "UnsafePointer"
	default:
		str += fmt.Sprint(strings.Repeat(" ", self.pc), v.Interface())
	}

	return str
}

//////////////////////////////////////////////////////////////////////
//url string
type EscapeError string

func (e EscapeError) Error() string {
	return "invalid URL escape " + strconv.Quote(string(e))
}

type encoding int

const (
	encodePath encoding = 1 + iota
	encodeQueryComponent
)

// QueryUnescape does the inverse transformation of QueryEscape, converting
// %AB into the byte 0xAB and '+' into ' ' (space). It returns an error if
// any % is not followed by two hexadecimal digits.
func QueryUnescape(s string) (string, error) {
	return unescape(s, encodeQueryComponent)
}

//query:k=v&k2=v2 ==>map
func ParseUrlParam2Map(m map[string]string, query string) *Error {
	var err error
	for query != "" {
		key := query
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		m[key] = value
	}
	if err != nil {
		return XErr(err)
	}
	return nil
}

func UnHex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func IsHex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unescape(s string, mode encoding) (string, error) {
	// Count %, check that they're well-formed.
	n := 0
	hasPlus := false
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			n++
			if i+2 >= len(s) || !IsHex(s[i+1]) || !IsHex(s[i+2]) {
				s = s[i:]
				if len(s) > 3 {
					s = s[:3]
				}
				return "", EscapeError(s)
			}
			i += 3
		case '+':
			hasPlus = mode == encodeQueryComponent
			i++
		default:
			i++
		}
	}

	if n == 0 && !hasPlus {
		return s, nil
	}

	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			t[j] = UnHex(s[i+1])<<4 | UnHex(s[i+2])
			j++
			i += 3
		case '+':
			if mode == encodeQueryComponent {
				t[j] = ' '
			} else {
				t[j] = '+'
			}
			j++
			i++
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t), nil
}

func EscapeSql(buf []byte, v string) []byte {
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
