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
