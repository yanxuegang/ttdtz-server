package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	models "juliang/model"
	"juliang/tool"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"unsafe"

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
	wxOrderQueryUrl = "https://api.mch.weixin.qq.com/pay/orderquery"
	AppID           = "wx7756c4bb1e711e79"
	AppSecret       = "d625e131569517cc8829cc99ed1154b9"
	MchId           = "1541123451"
)

var (
	DB    = &gorm.DB{}
	chuan = "b0ce784009590860e2a62fa5deaed8f0"
	kuai  = "6XD6wD7IYozyu8hHESWSRdriznEnvS6t"
	you   = "4b178983b3ccd0cc7c2b9cf64fab714d"
	csj   = 0
	ks    = 1
	ylh   = 2
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
	translation := models.Translations{
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
		//Ua:        req.Form.Get("ua"),
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
		Translation = new(models.Translations)
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

	//if time.Now().Unix()-604800 > Translation.Ts.Unix() {
	//DB.Where("oaid_md5 = ?", Translation.OaidMd5).Delete(Translations{})
	// if err != nil {
	// 	log.Println("gameServerHandler: ", err)
	// }
	//log.Printf("check fail timeout md5 = %s", oaid_md5)
	//msg, _ := json.Marshal(gameResponse{Code: 1, Msg: "timeout"})
	//w.Write(msg)
	//return
	//}
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
		m = tool.GetMd5String1(transId, secKey)
	default:
		m = tool.GetMd5String1(secKey, transId)
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

func wxpaynoticeHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("wxpaynoticeHandler req")
	err := req.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	formData := make(map[string]interface{})
	json.NewDecoder(req.Body).Decode(&formData)
	log.Println(req.Form)
	total_fee, _ := strconv.Atoi(req.Form.Get("total_fee"))
	settlement_total_fee, _ := strconv.Atoi(req.Form.Get("settlement_total_fee"))
	cash_fee, _ := strconv.Atoi(req.Form.Get("cash_fee"))
	time_end, _ := time.Parse("2006-01-02", req.Form.Get("time_end"))
	order := models.Orders{
		Appid:              req.Form.Get("appid"),
		MchId:              req.Form.Get("mch_id"),
		Openid:             req.Form.Get("openid"),
		TotalFee:           float64(total_fee),
		SettlementTotalFee: settlement_total_fee,
		FeeType:            req.Form.Get("fee_type"),
		CashFee:            cash_fee,
		TransactionId:      req.Form.Get("transaction_id"),
		OutTradeNo:         req.Form.Get("out_trade_no"),
		TimeEnd:            time_end,
	}
	if err = order.Create(); err != nil {
		//if err != nil {
		log.Println("Create order err: ", err)
	}
	w.Write([]byte("SUCCESS"))
}

func wxpaymentHandler(w http.ResponseWriter, req *http.Request) {
	//check local mysql
	var (
		data = map[string]interface{}{
			"appid":        AppID,
			"mch_id":       MchId,
			"out_trade_no": "",
			"nonce_str":    tool.RandString(10),
		}
	)
}

func main() {
	sendhttpserver()
	var err error
	DB, err = models.NewDBEngine()
	if err != nil {
		log.Fatal("NewDBEngine: ", err)
	}
	log.Printf("DBs init %+v", DB)
	//
	http.HandleFunc("/adcallback", listenAdHandler)
	http.HandleFunc("/apicallback", gameServerHandler)
	//
	http.HandleFunc("/v1/chuan-notice", videochuanHandler)
	http.HandleFunc("/v1/kuai-notice", videokuanHandler)
	http.HandleFunc("/v1/you-notice", videoyouHandler)
	//
	http.HandleFunc("/v1/wxpay-notice", wxpaynoticeHandler)
	http.HandleFunc("/v1/wxpayment", wxpaymentHandler)
	//http.ListenAndServeTLS(":8081", "server.crt",
	//	"server.key", nil)
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func sendhttpserver() {
	data := make(map[string]interface{})
	params := make(map[string]interface{})
	params["open_id"] = "1"
	params["type"] = "1"
	params["password"] = "1"
	params["channel"] = "1"
	data["cmd"] = 1000
	data["params"] = params
	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://192.168.1.6:8000/api/v1/login"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}
