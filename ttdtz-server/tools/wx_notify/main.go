package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

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

func main() {
	http.HandleFunc("/v1/wxpay-notice", wxpaynoticeHandler)
	http.ListenAndServeTLS(":8082", "server.crt",
		"server.key", nil)
}
