package login_test

import (
	"net/http"
	"testing"
)

func Login_test(t *testing.T) {
	url := "http://192.168.1.6:8000/api/v1/login"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Log(err)
		//return nil, errors.New("new request is fail ")
	}

	q := req.URL.Query()
	q.Add("cmd", "1000") //"EJiw267wvfQCGKf2g74ZIPD89-vIATAMOAFCIjIwMTkxMTI3MTQxMTEzMDEwMDI2MDc3MjE1MTUwNTczNTBIAQ==")
	//q.Add("params", params.Get("imei")) //"0c2bd03c39f19845bf54ea0abafae70e")
	//q.Add("event_type", "0")
	//q.Add("conv_time", strconv.Itoa(int(time.Now().Unix())))
	//log.Println("translationa Q = ", q)
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	client.Do(req)
}
