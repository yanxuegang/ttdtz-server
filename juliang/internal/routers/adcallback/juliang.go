package adcallback

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func JuliangAdHandler(w http.ResponseWriter, req *http.Request) {
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
