// client.go
package main

import (
	"fmt"
	"os"

	"../common"
	//"time"
)


func analysis() {

}


func invoicing() {
	var err error
	var r string


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
