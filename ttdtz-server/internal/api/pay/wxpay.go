package pay

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"ttdtz-server/global"
	"ttdtz-server/internal/managers"
	"ttdtz-server/pkg/app"
	"ttdtz-server/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Conf struct {
	Pay `xml:"xml"`
}

type Pay struct {
	mch_appid        string `xml:"mch_appid"`
	mchid            string `xml:"mchid"`
	nonce_str        string `xml:"nonce_str"`
	openid           string `xml:"openid"`
	check_name       string `xml:"check_name"`
	re_user_name     string `xml:"re_user_name"`
	amount           string `xml:"amount"`
	partner_trade_no string `xml:"partner_trade_no"`
	desc             string `xml:"desc"`
	spbill_create_ip string `xml:"spbill_create_ip"`
	sign             string `xml:"sign"`
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

func WxPay(c *gin.Context) {
	req := managers.WxPayRequestInfo{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	log.Printf("WxPay req => %+v", req)

	c.String(http.StatusOK, "")
	respdata, err := managers.WxPay(c.Request.Context(), &req)
	if err != nil {
		global.Logger.Error(err)
		response.ToErrorResponse(errcode.InvalidParams)
		//return
	}
	response.ToResponse(respdata)
}
