package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	models "juliang/model"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
)

//request
/*IMEI          安卓的设备 ID 的 md5 摘要，32位
  IDFA          IOS 6+的设备id字段，32
  OAID          Android Q及更高版本的设备号，32位
  OAID_MD5      Android Q及更高版本的设备号的md5摘要，32位
  CAID1、CAID2  不同版本的中国广告协会互联网广告标识，CAID1是20201230版，暂无CAID2
  OS            操作系统平台
*/

//response
/*callback 是（点击事件）点击检测下发的 callback EJiw267wvfQCGKf2g74ZIPD89-vIATAMOAFCIjIwMTkxMTI3MTQxMTEzMDEwMDI2MDc3MjE1MTUwNTczNTBIAQ==
  imei 是  安卓手机 imei 的 md5 摘要 0c2bd03c39f19845bf54ea0abafae70e
  idfa 是  ios 手机的 idfa 原值 FCD369C3-F622-44B8-AFDE-12065659F34B
  muid 是  安卓：imei号取md5sum摘要；IOS：取idfa原值 FCD369C3-F622-44B8-AFDE-12065659F34B
  oaid 是  Android Q 版本的 oaid 原值 b305ee2fefddfea2
  oaid_md5 否 Android Q 版本的 oaid 原值的md5摘要 8FCF82C6-47E7-2679-2F44-37405B982580
  caid1 、caid2 是 不同版本版本的中国广告协会互联网广告标识，caid1为最新版本，caid2为老版本 f949f306494646edfee1f939698e1fb1
  os 是 客户端的操作系统类型0: android 1: ios
  source 否 数据来源，广告主可自行定义
  conv_time 否 转化发生的时间，UTC 时间戳 1574758519
  event_type 是 事件类型见列表
  match_type 否 归因方式 0:点击 1:展示 2:有效播放归因*/

var (
	translationaUrl = "https://ad.oceanengine.com/track/activate/"
)

//游戏api上报数据请求
type GameRequestData struct {
	Imei      string `json:"imei"`
	Idfa      string `json:"idfa"`
	Muid      string `json:"muid"`
	Oaid      string `json:"oaid"`
	OaidMd5   string `json:"oaid_md5"`
	Os        int    `json:"os"`
	EventType int    `json:"event_type"`
	AndroidId string `json:"android_id"`
	Mac       string `json:"mac"`
	Mac1      string `json:"mac1"`
	Ip        string `json:"ip"`
	Ua        string `json:"ua"`
	Ts        string `json:"ts"`
	Model     string `json:"model"`
}

//转化事件回调返回值
type TrResponseData struct {
	Code int    `json:"code"`
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
}

type VideoResult struct {
	IsValid bool `json:"isValid"`
}

type gameResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//转化事件回调请求
type TrRequestData struct {
}

//监测数据sqlmodel
type Translations struct {
	Id          int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Aid         string    `gorm:"column:aid" db:"aid" json:"aid" form:"aid"`
	ConvertId   string    `gorm:"column:convert_id" db:"convert_id" json:"convert_id" form:"convert_id"`
	RequestId   string    `gorm:"column:request_id" db:"request_id" json:"request_id" form:"request_id"`
	Imei        string    `gorm:"column:imei" db:"imei" json:"imei" form:"imei"`
	Idfa        string    `gorm:"column:idfa" db:"idfa" json:"idfa" form:"idfa"`
	Androidid   string    `gorm:"column:androidid" db:"androidid" json:"androidid" form:"androidid"`
	Oaid        string    `gorm:"column:oaid" db:"oaid" json:"oaid" form:"oaid"`
	OaidMd5     string    `gorm:"column:oaid_md5" db:"oaid_md5" json:"oaid_md5" form:"oaid_md5"`
	Os          int       `gorm:"column:os" db:"os" json:"os" form:"os"`
	Mac         string    `gorm:"column:mac" db:"mac" json:"mac" form:"mac"`
	Mac1        string    `gorm:"column:mac1" db:"mac1" json:"mac1" form:"mac1"`
	Ip          string    `gorm:"column:ip" db:"ip" json:"ip" form:"ip"`
	Ua          string    `gorm:"column:ua" db:"ua" json:"ua" form:"ua"`
	Geo         string    `gorm:"column:geo" db:"geo" json:"geo" form:"geo"`
	Ts          time.Time `gorm:"column:ts" db:"ts" json:"ts" form:"ts"`
	CallbackUrl string    `gorm:"column:callback_url" db:"callback_url" json:"callback_url" form:"callback_url"`
	Callback    string    `gorm:"column:callback" db:"callback" json:"callback" form:"callback"`
	Model       string    `gorm:"column:model" db:"model" json:"model" form:"model"`
	Status      int       `gorm:"column:status" db:"status" json:"status" form:"status"`
}

func (m *Translations) Create() error {
	//todo redis save
	return DB.Create(m).Error
}

func (m *Translations) Delete() error {
	//todo redis save
	return DB.Delete(m).Error
}

func (m *Translations) Update(attrs ...interface{}) error {
	//todo redis save
	rowsAffected := DB.Model(m).Update(attrs...).RowsAffected
	if rowsAffected == 0 {
		log.Println("[PLAYER][WARNING] No Content To Update.")
		return nil
	}
	return nil
}

func listenAdHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("listenAdHandler req")
	err := req.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(req.Body).Decode(&formData)
	log.Println(req.Form)

	w.Write([]byte("success"))

	resp, err := translationa(req.Form)
	if err != nil {
		log.Println("translationa: ", err)
		return
	}
	var (
		wxrespData TrResponseData
		json       = jsoniter.ConfigCompatibleWithStandardLibrary
	)

	if err2 := json.NewDecoder(resp.Body).Decode(&wxrespData); err2 != nil {
		log.Println("resp Decode: ", err2)
		return
	}
	log.Println("translationa data : ", wxrespData)
	//translation := new(Translations)
	os, _ := strconv.Atoi(req.Form.Get("os"))
	translation := Translations{
		Imei: req.Form.Get("imei"),
		Idfa: req.Form.Get("idfa"),
		//Muid:    req.Form.Get("muid"),
		Oaid:    req.Form.Get("oaid"),
		OaidMd5: req.Form.Get("oaid_md5"),
		Ts:      time.Now(),
		Os:      os,
		// 	//EventType: req.Form.Get("event_type"),
		Androidid: req.Form.Get("android_id"),
		Mac:       req.Form.Get("mac"),
		Mac1:      req.Form.Get("mac1"),
		Ip:        req.Form.Get("ip"),
		Ua:        req.Form.Get("ua"),

		// 	Model: req.Form.Get("model"),
	}
	//err = DBs["app_line"].Exec("INSERT INTO `translations` VALUES (?, ?, ?)", translation.Imei, translation.Idfa, "").Error
	if err = translation.Create(); err != nil {
		//if err != nil {
		log.Println("Create Translation err: ", err)
	}

}

func translationa(params url.Values) (*http.Response, error) {
	url := "https://ad.oceanengine.com/track/activate/"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
	for key, value := range params {
		log.Println("translationa key:", key, " => value :", value)
	}
	q := req.URL.Query()
	q.Add("callback", params.Get("callback_param")) //"EJiw267wvfQCGKf2g74ZIPD89-vIATAMOAFCIjIwMTkxMTI3MTQxMTEzMDEwMDI2MDc3MjE1MTUwNTczNTBIAQ==")
	q.Add("imei", params.Get("imei"))               //"0c2bd03c39f19845bf54ea0abafae70e")
	q.Add("event_type", "0")
	q.Add("conv_time", strconv.Itoa(int(time.Now().Unix())))
	log.Println("translationa Q = ", q)
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	return client.Do(req)
}

func gameServerHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("gameServerHandler req")

	err := req.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	for k, v := range req.Header {
		log.Println("Header.key:", k, " => Header.value :", v)
	}
	formData := make(map[string]interface{})
	json.NewDecoder(req.Body).Decode(&formData)
	log.Println(req.Form)
	for key, value := range req.Form {
		log.Println("key:", key, " => value :", value)
	}
	oaid_md5 := req.Form.Get("OaidMd5")
	w.Header().Set("content-type", "text/json")
	var (
		Translation = new(Translations)
	)
	Translation.OaidMd5 = oaid_md5
	log.Printf("oaid_md5 = %s, Translation.OaidMd5 = %s", oaid_md5, Translation.OaidMd5)
	err = DB.Where("oaid_md5 = ?", oaid_md5).Find(&Translation).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("gameServerHandler: ", err)
		log.Printf("check fail ErrRecordNotFound md5 = %s", oaid_md5)
		msg, _ := json.Marshal(gameResponse{Code: 1, Msg: "NotFound"})
		w.Write(msg)
		return
	} else if err == gorm.ErrRecordNotFound {
		// MySQL未查询到
		log.Printf("check fail NotFound md5 = %s", oaid_md5)
		msg, _ := json.Marshal(gameResponse{Code: 1, Msg: "NotFound"})
		w.Write(msg)
		return
	}

	if time.Now().Unix()-604800 > Translation.Ts.Unix() {
		//DB.Where("oaid_md5 = ?", Translation.OaidMd5).Delete(Translations{})
		// if err != nil {
		// 	log.Println("gameServerHandler: ", err)
		// }
		log.Printf("check fail timeout md5 = %s", oaid_md5)
		//msg, _ := json.Marshal(gameResponse{Code: 1, Msg: "timeout"})
		//w.Write(msg)
		return
	}
	//DB.Where("oaid_md5 = ?", Translation.OaidMd5).Delete(Translations{})
	// err = Translation.Delete()
	// if err != nil {
	// 	log.Println("gameServerHandler: ", err)
	// }
	log.Printf("check success md5 = %s", oaid_md5)
	msg, _ := json.Marshal(gameResponse{Code: 0, Msg: "success"})
	w.Write(msg)
	//w.Write([]byte("success"))
}

var (
	DB    = &gorm.DB{}
	chuan = "b4e2aa01a2d47c36deb2d24d8552acc8"
	kuai  = "4b178983b3ccd0cc7c2b9cf64fab714d"
	you   = "WKzX2uhtpDEIYyVixkHWhNpif3edFCpW"
	csj   = 0
	ks    = 1
	ylh   = 2
)

func getMd5String1(str1, str2 string) string {
	md5str := str1 + ":" + str2
	m := md5.New()
	_, err := io.WriteString(m, md5str)
	if err != nil {
		log.Fatal("getMd5String1 error ", err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func videoyouHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("videoyouHandler req")
	err := req.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(req.Body).Decode(&formData)
	log.Println(req.Form)
	resp, _ := do(req, you, ylh)

	w.Header().Set("content-type", "text/json")
	msg, _ := json.Marshal(resp)
	w.Write(msg)
}

func videokuanHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("videokuanHandler req")
	err := req.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(req.Body).Decode(&formData)
	log.Println(req.Form)
	resp, _ := do(req, kuai, ks)

	w.Header().Set("content-type", "text/json")
	msg, _ := json.Marshal(resp)
	w.Write(msg)
}

func do(req *http.Request, secKey string, Type int) (VideoResult, bool) {
	sign := req.Form.Get("sign")
	transId := req.Form.Get("transid")
	//extra := req.Form.Get("extra")
	var m string
	switch Type {
	case ylh:
		m = getMd5String1(transId, secKey)
	default:
		m = getMd5String1(secKey, transId)
	}

	log.Printf("sign = %s, Md5 = %s", sign, m)
	if sign != m {
		return VideoResult{IsValid: false}, true
		//log.Fatal("sign Md5 error ", err)
	} else {
		return VideoResult{IsValid: true}, true
	}
}

func videochuanHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("videochuanHandler req")
	err := req.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	formData := make(map[string]interface{})
	json.NewDecoder(req.Body).Decode(&formData)
	log.Println(req.Form)
	resp, _ := do(req, chuan, csj)

	w.Header().Set("content-type", "text/json")
	msg, _ := json.Marshal(resp)
	w.Write(msg)
	//w.Write([]byte("videomonitoringHandler success"))
}

func main() {
	var err error
	DB, err = models.NewDBEngine()
	if err != nil {
		log.Fatal("NewDBEngine: ", err)
	}
	log.Printf("DBs init %+v", DB)

	http.HandleFunc("/adcallback", listenAdHandler)
	http.HandleFunc("/apicallback", gameServerHandler)
	http.HandleFunc("/v1/chuan-notice", videochuanHandler)
	http.HandleFunc("/v1/kuai-notice", videokuanHandler)
	http.HandleFunc("/v1/you-notice", videoyouHandler)
	http.ListenAndServeTLS(":8088", "server.crt",
		"server.key", nil)
	// err = http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// 	return
	// }
}
