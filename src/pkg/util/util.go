package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ParseAccessToken(rawToken string) ([]byte, string, error) {
	token := strings.Split(rawToken, " ")
	if len(token) <= 1 {
		return nil, "", errors.New("Invalid bearer token not valid")
	}
	if !InArray(strings.ToLower(token[0]), []string{"bearer", "basic"}) {
		return nil, "", errors.New("Invalid token type")
	}

	tokenPayloads := strings.Split(token[1], ".")
	if len(tokenPayloads) <= 1 {
		return nil, "", errors.New("Invalid bearer token not valid")
	}

	payloadIssue, err := base64.RawStdEncoding.DecodeString(tokenPayloads[1])
	if err != nil {
		return nil, "", errors.New("Invalid issued profile token")
	}
	return payloadIssue, token[1], nil
}

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int:
		return strconv.FormatInt(int64(value), 10)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}

// StringToInt ..
func StringToInt(str string) int {
	i1, _ := strconv.Atoi(str)
	return i1
}

// StringToInt64 ...
func StringToInt64(str string) int64 {
	i64, _ := strconv.ParseInt(str, 10, 64)
	return i64
}

// StrToUint64 ...
func StrToUint64(value string) uint64 {
	res, _ := strconv.ParseUint(string(value), 10, 64)
	return res
}

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

func (b *buffer) indent() {
	if len(b.r) > 0 {
		b.r = append(b.r, '_')
	}
}

// ToSnackeCase ...
func ToSnackeCase(s string) string {
	b := buffer{
		r: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.indent()
					w = true
				}
				b.write(m)
			}
			m = unicode.ToLower(ch)
		} else {
			if m != 0 {
				b.indent()
				b.write(m)
				m = 0
				w = false
			}
			b.write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.indent()
		}
		b.write(m)
	}
	return string(b.r)
}
