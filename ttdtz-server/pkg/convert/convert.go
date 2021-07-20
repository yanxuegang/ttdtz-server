package convert

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
	"ttdtz-server/global"
)

type StrTo string

func (s StrTo) String() string {
	return string(s)
}
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

func stringifyData(data map[string]interface{}) string {
	var (
		keys []string
		vals []string
	)
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		vals = append(vals, fmt.Sprintf("%s=%v", k, data[k]))
	}

	return strings.Join(vals, "&")
}

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func GetWxpayMd5(data map[string]interface{}) string {
	message := stringifyData(data)
	message = message + "&key=" + global.GlobalConfig.Wx.AppSecret
	log.Println("GetWxpayMd5", message)

	m := md5.New()
	_, err := io.WriteString(m, message)
	if err != nil {
		log.Fatal("getMd5String1 error ", err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
