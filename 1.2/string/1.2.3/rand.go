package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// 获取随机字符串
func GetRandomString(count int) string {
	// 声明字符串
	str := "0123456789avcdefghijklmnopqrstuvwxyz"
	// 转换数组
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 使用模块crypto/rand 生成随机字符串
func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length forNewLenchars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // 存储随机字符串
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // 跳过此数字避免偏差
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}

}

// 用crypto/md5 实现随机密码
func CreatePassword() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "shirdon.liaosdfsf@#")
	io.WriteString(h, t.String())
	password := fmt.Sprintf("%x", h.Sum(nil))
	return password
}
func main() {
	// 获取10位随机字符串
	b := GetRandomString(10)
	fmt.Println(b)
	var data []byte = []byte("0123456789avcdefghijlmnopqrstuvwxyz")
	c := NewLenChars(20, data)
	fmt.Println(c)

	pass := CreatePassword()
	fmt.Println(pass)
}
