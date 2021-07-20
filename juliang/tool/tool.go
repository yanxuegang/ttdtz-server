package tool

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func GetWxpayMd5(appid, mch_id, out_trade_no, nonce_str, key string) string {
	message := stringifyData(data)
	message += fmt.Sprintf("&org_loc=%s&method=POST&secret=%s", uri, GetMidasPaySecret())
	log.Println("getSig", message)
	return ""
}

func GetMd5String1(str1, str2 string) string {
	md5str := str1 + ":" + str2
	m := md5.New()
	_, err := io.WriteString(m, md5str)
	if err != nil {
		log.Fatal("getMd5String1 error ", err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
