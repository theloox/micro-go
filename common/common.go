// common.go
package common

import (
	"fmt"
	"io"
	"os"
	"strings"
	"io/ioutil"
	"net/http"
)


const HOST string = "http://localhost"

const A_PORT = 12001
const I_PORT = 12002
const R_PORT = 12003


type Invoice struct {
	_id int64		`json:"_id,omitempty"`
	Id int64			`json:"id,omitempty"`
	Uid uint32		`json:"uid,omitempty"`
	Utime int64		`json:"utime,omitempty"`
	Client float64	`json:"client,omitempty"`
	Items float64	`json:"items,omitempty"`
	Amount float64	`json:"amount,omitempty"`
}


func Xhttp(m string, p int, u string, b string) (string, error) {
	var cl *http.Client
	var req *http.Request
	var res *http.Response
	var r []byte
	var rd io.Reader
	var err error

	var ret string


	cl = &http.Client{}

	url := fmt.Sprintf("%s:%d%s", HOST, p, u)

	if (b == "") {
		req, err = http.NewRequest(m, url, nil)
	} else {
		rd = strings.NewReader(b)
		req, err = http.NewRequest(m, url, rd)
		if (err != nil) {
			fmt.Fprintf(os.Stderr, "xhttp: %v\n\n", err)
			return ret, err
		}
	}


	req.Header.Add("Connection", "close")
	req.TransferEncoding = []string{"text/json"}

	res, err = cl.Do(req)
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "xhttp: %v\n\n", err)
		return ret, err
	}
	defer res.Body.Close()


	r, err = ioutil.ReadAll(res.Body)
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "xhttp: %v\n\n", err)
		return ret, err
	}

	ret = string(r)

	return ret, err
}
