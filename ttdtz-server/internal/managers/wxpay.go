package managers

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"ttdtz-server/global"
	"ttdtz-server/pkg/convert"
	"unsafe"
)

var privateKey, publicKey []byte

func init() {
	var err error
	publicKey, err = ioutil.ReadFile("/home/yxg1994/workspace/src/ttdtz-server/apiclient_key.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("/home/yxg1994/workspace/src/ttdtz-server/apiclient_cert.pem")
	if err != nil {
		os.Exit(-1)
	}
}

type payparams struct {
	OpenId        string `json:"open_id" form:"open_id"`
	Type          string `json:"type" form:"type"`
	TransactionId string `json:"transaction_id" form:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no" form:"out_trade_no"`
	Total         int    `json:"total" form:"total"`
	Channel       string `json:"channel" form:"channel"`
}

type WxPayRequestInfo struct {
	Cmd    int       `json:"cmd" form:"cmd"`
	Params payparams `json:"params" form:"params"`
}

type WxPayResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func BuildErrorWxPayResponse(code int, msg string) *WxPayResponse {
	response := &WxPayResponse{}
	response.Code = code
	response.Msg = msg
	return response
}

func BuildWxPayResponse() *WxPayResponse {
	response := &WxPayResponse{}
	response.Code = 0
	return response
}

type Conf struct {
	Pay `xml:"xml"`
}

type Pay struct {
	Mch_appid  string `xml:"mch_appid"`
	Mchid      string `xml:"mchid"`
	Nonce_str  string `xml:"nonce_str"`
	Openid     string `xml:"openid"`
	Check_name string `xml:"check_name"`
	//Re_user_name     string `xml:"re_user_name"`
	Amount           int    `xml:"amount"`
	Partner_trade_no string `xml:"partner_trade_no"`
	Desc             string `xml:"desc"`
	//Spbill_create_ip string `xml:"spbill_create_ip"`
	Sign string `xml:"sign"`
}

func ReadFile() Pay {
	file_data, err := file_open("./wxpay.xml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	tpl := Pay{}
	//根据传递的构造体解析 切记 第二个参数传递的是指针
	tpl_err := xml.Unmarshal(file_data, &tpl)
	if tpl_err != nil {
		fmt.Println("tpl err :", tpl_err)
		os.Exit(0)
	}
	fmt.Println(tpl)
	return tpl
}

func file_open(file_name string) ([]byte, error) {
	//定义变量
	var (
		open               *os.File
		file_data          []byte
		open_err, read_err error
	)
	//打开文件
	open, open_err = os.Open(file_name)
	if open_err != nil {
		return nil, open_err
	}
	//关闭资源
	defer open.Close()

	//读取所有文件内容
	file_data, read_err = ioutil.ReadAll(open)
	if read_err != nil {
		return nil, read_err
	}
	return file_data, nil
}

func WxPay(ctx context.Context, req *WxPayRequestInfo) (*WxPayResponse, error) {
	var (
		mapData = map[string]interface{}{
			"mch_appid":        global.GlobalConfig.Wx.AppId,
			"mchid":            global.GlobalConfig.Wx.MchId,
			"nonce_str":        convert.RandString(10),
			"partner_trade_no": convert.RandString(11),
			"openid":           req.Params.OpenId,
			"check_name":       "NO_CHECK",
			"amount":           int(1),
			"desc":             "游戏",
		}
	)
	var data Pay
	data.Mch_appid = global.GlobalConfig.Wx.AppId
	data.Mchid = global.GlobalConfig.Wx.MchId
	data.Nonce_str = mapData["nonce_str"].(string)
	data.Partner_trade_no = mapData["partner_trade_no"].(string)
	data.Openid = req.Params.OpenId
	data.Check_name = "NO_CHECK"
	data.Amount = int(1)
	data.Desc = "游戏"
	sign := convert.GetWxpayMd5(mapData)

	log.Println("GetWxpayMd5", sign)
	data.Sign = sign

	// str1 := "amount=" + "1&" + "check_name=" + "NO_CHECK&" + "desc=" + "游戏&" + "mch_appid=" + "wx7756c4bb1e711e79&" + "mchid=" + "1541123451&" + "nonce_str=" + "MZVQBIBBDF&" + "openid=" + "o3J8n6-s5YRWURvRNoXWSqyUXGgA&" + "partner_trade_no=" + "MZVQBIBBDFH&" + "key=" + "d625e131569517cc8829cc99ed1154b9"
	// m := md5.New()
	// _, err := io.WriteString(m, str1)
	// if err != nil {
	// 	log.Fatal("getMd5String1 error ", err)
	// }
	// arr := m.Sum(nil)
	// data.Sign = strings.ToTitle(fmt.Sprintf("%x", arr))

	bytexml, err := xml.MarshalIndent(&data, "", "  ")
	fmt.Println(string(bytexml))
	c, _ := tls.X509KeyPair(privateKey, publicKey)
	cfg := &tls.Config{
		Certificates: []tls.Certificate{c},
	}
	tr := &http.Transport{
		TLSClientConfig: cfg,
	}
	client := &http.Client{Transport: tr}
	wxreq, err := http.NewRequest("POST", global.GlobalConfig.Wx.WxPay, bytes.NewBuffer(bytexml))
	if err != nil {
		fmt.Println(err)
	}
	wxreq.Header.Add("Content-Type", "application/xml; charset=utf-8")

	resp, err := client.Do(wxreq)
	if err != nil {
		return BuildErrorWxPayResponse(1, err.Error()), err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	// var respData map[string]interface{}
	// if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {
	// 	return BuildErrorWxPayResponse(2, err.Error()), err
	// }
	// log.Printf("WxPay respData = %+v", respData)
	// if _, ok := respData["return_code"].(string); !ok {
	// 	return BuildErrorWxPayResponse(3, "WxPay FALL"), nil
	// }
	// if respData["return_code"].(string) != "SUCCESS" {
	// 	return BuildErrorWxPayResponse(4, "WxPay FALL"), nil
	// }
	return BuildWxPayResponse(), nil
}
