// client.go
package main

import (
	"fmt"
	"os"
	"encoding/json"


	"../common"
	//"time"
)


func analysis() {
	var err error
	var r string
	var buf []byte


	err = err


	fmt.Printf("================\n")


	fmt.Printf("Testing root\n")
	r, _ = common.Xhttp("GET", common.A_PORT, "/", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing root POST\n")
	r, _ = common.Xhttp("POST", common.A_PORT, "/", "qwerty")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


    fmt.Printf("Testing abnormal client #1\n");
	r, _ = common.Xhttp("GET", common.A_PORT, "/abnormal/1", "")
	fmt.Printf("%s\n", r)

	fmt.Printf("Testing create\n")
	buf, _ = json.Marshal(common.Invoice{Client: 1, Items:10, Amount: 4200.00})
	r, _ = common.Xhttp("POST", common.I_PORT, "/create", string(buf))
	fmt.Printf("%s\n", r)

    fmt.Printf("Testing abnormal client #1\n");
	r, _ = common.Xhttp("GET", common.A_PORT, "/abnormal/1", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


    fmt.Printf("Testing duplicates\n");
	r, _ = common.Xhttp("GET", common.A_PORT, "/duplicates", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


    fmt.Printf("Testing same uid\n");
	r, _ = common.Xhttp("GET", common.A_PORT, "/same", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing create\n")
	buf, _ = json.Marshal(common.Invoice{Client: 3, Items:33, Amount: 330.00, Uid: 33333333})
	r, _ = common.Xhttp("POST", common.I_PORT, "/create", string(buf))
	r, _ = common.Xhttp("POST", common.I_PORT, "/create", string(buf))
	fmt.Printf("%s\n", r)

    fmt.Printf("Testing same uid\n");
	r, _ = common.Xhttp("GET", common.A_PORT, "/same", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


}


func invoicing() {
	var err error
	var r string
	var buf []byte


	err = err


	fmt.Printf("================\n")


	fmt.Printf("Testing root\n")
	r, _ = common.Xhttp("GET", common.I_PORT, "/", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing root POST\n")
	r, _ = common.Xhttp("POST", common.I_PORT, "/", "qwerty")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


	fmt.Printf("Testing all\n")
	r, _ = common.Xhttp("GET", common.I_PORT, "/all", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


	fmt.Printf("Testing create\n")
	buf, _ = json.Marshal(common.Invoice{Client: 1, Items:10, Amount: 420.00})
	r, _ = common.Xhttp("POST", common.I_PORT, "/create", string(buf))
	fmt.Printf("%s\n", r)

	buf, _ = json.Marshal(common.Invoice{Client: 2, Items: 2, Amount: 69.00})
	r, _ = common.Xhttp("POST", common.I_PORT, "/create", string(buf))
	fmt.Printf("%s\n", r)

	buf, _ = json.Marshal(common.Invoice{Client: 3, Items: 1, Amount: 9.00})
	r, _ = common.Xhttp("POST", common.I_PORT, "/create", string(buf))
	fmt.Printf("%s\n", r)

	fmt.Printf("================\n")

	fmt.Printf("Testing all\n")
	r, _ = common.Xhttp("GET", common.I_PORT, "/all", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


	fmt.Printf("Testing read #1\n")
	r, _ = common.Xhttp("GET", common.I_PORT, "/read/1", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing read #9999 (not found)\n")
	r, _ = common.Xhttp("GET", common.I_PORT, "/read/9999", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


	fmt.Printf("Testing update #3 items 1->3\n")
	buf, _ = json.Marshal(common.Invoice{Id: 3, Client: 3, Items:3, Amount: 27.00})
	r, _ = common.Xhttp("PUT", common.I_PORT, "/update", string(buf))
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing update #9999 (not found)\n")
	buf, _ = json.Marshal(common.Invoice{Id: 9999, Client: 3, Items:3, Amount: 27.00})
	r, _ = common.Xhttp("PUT", common.I_PORT, "/update", string(buf))
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing all\n")
	r, _ = common.Xhttp("GET", common.I_PORT, "/all", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


	fmt.Printf("Testing delete #2\n")
	r, _ = common.Xhttp("DELETE", common.I_PORT, "/delete/2", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing delete #9999 (not found)\n")
	r, _ = common.Xhttp("DELETE", common.I_PORT, "/delete/9999", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing all\n")
	r, _ = common.Xhttp("GET", common.I_PORT, "/all", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


}


func reporting() {
	var err error
	var r string


	err = err


	fmt.Printf("================\n")
	fmt.Printf("p %d\n", common.R_PORT)


	fmt.Printf("Testing root\n")
	r, _ = common.Xhttp("GET", common.R_PORT, "/", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

	fmt.Printf("Testing root POST\n")
	r, _ = common.Xhttp("POST", common.R_PORT, "/", "qwerty")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


    fmt.Printf("Testing all\n");
	r, _ = common.Xhttp("GET", common.R_PORT, "/all", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


    fmt.Printf("Testing client(field) #1\n");
	r, _ = common.Xhttp("GET", common.R_PORT, "/client/1", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

    fmt.Printf("Testing client(field) #9999 (not found)\n");
	r, _ = common.Xhttp("GET", common.R_PORT, "/client/9999", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


    fmt.Printf("Testing last 60 (seconds)\n");
	r, _ = common.Xhttp("GET", common.R_PORT, "/last/60", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")

    fmt.Printf("Testing last 3600 (seconds)\n");
	r, _ = common.Xhttp("GET", common.R_PORT, "/last/3600", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


    fmt.Printf("Testing today\n");
	r, _ = common.Xhttp("GET", common.R_PORT, "/today", "")
	fmt.Printf("%s\n", r)
	fmt.Printf("================\n")


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
