// client.go
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"io/ioutil"
	"net/http"

	"../common"
	//"time"
)


const HOST string = "http://localhost"


func xhttp(m string, p int, u string, b string) (string, error) {
	var cl *http.Client
	var req *http.Request
	var res *http.Response
	var r []byte
	var rd io.Reader
	var err error

	var ret string


	cl = &http.Client{}

	url := fmt.Sprintf("%s:%d%s", HOST, common.I_PORT, u)

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


func analysis() {

}


func invoicing() {
	var err error
	var r string


	err = err


	fmt.Printf("================\n")


	fmt.Printf("Testing root\n")
	r, _ = xhttp("GET", common.I_PORT, "/", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing root POST\n")
	r, _ = xhttp("POST", common.I_PORT, "/", "qwerty")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


}


func reporting() {

}


func main() {

	if (len(os.Args) < 2) {
		fmt.Fprintf(os.Stderr, "Usage: xclient analysis|invoicing|reporting\n\n");
		os.Exit(1);
	}

	if (os.Args[1] == "analysis") {
		analysis()
	} else if (os.Args[1] == "invoicing") {
		invoicing()
	} else if (os.Args[1] == "reporting") {
		reporting()
	} else {
		fmt.Printf("Usage: xclient analysis|invoicing|reporting\n");
	}

}
