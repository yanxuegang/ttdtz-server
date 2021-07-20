package managers

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"ttdtz-server/global"
	"ttdtz-server/pkg/convert"

	jsoniter "github.com/json-iterator/go"
)

type payparams struct {
	OpenId        string `json:"open_id" form:"open_id"`
	Type          string `json:"type" form:"type"`
	TransactionId string `json:"transaction_id" form:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no" form:"out_trade_no"`
	Channel       string `json:"channel" form:"channel"`
}

type WxPayRequestInfo struct {
	Cmd    int       `json:"cmd" form:"cmd"`
	Params payparams `json:"params" form:"params"`
}

type WxPayResponse struct {
	Code int `json:"code"`
}

func WxPay(ctx context.Context, req *WxPayRequestInfo) (*WxPayResponse, error) {
	var (
		data = map[string]interface{}{
			"appid":        global.GlobalConfig.Wx.AppId,
			"mch_id":       global.GlobalConfig.Wx.MchId,
			"out_trade_no": req.Params.OutTradeNo,
			"nonce_str":    convert.RandString(10),
		}
	)
	data["sign"] = convert.GetWxpayMd5(data)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonData, _ := json.Marshal(data)
	wxreq, err := http.NewRequest("POST", global.GlobalConfig.Wx.WxPayChechUrl, bytes.NewBuffer(jsonData))
	wxreq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(wxreq)
	if err != nil {

	}
	defer resp.Body.Close()

	var respData map[string]interface{}

	if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {

	}
	log.Println("response Status:", resp.Status)
	log.Println("response Body:", respData)
	return nil, nil
}
