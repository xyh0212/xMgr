package x

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
)

func Md5FromIO(r io.Reader) []byte {
	md5h := md5.New()
	io.Copy(md5h, r)
	b := md5h.Sum(nil)
	return b
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//生成随机字符串
func RandStr(nLen int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < nLen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
