package main

import (
	"fmt"
	"os"
	//"strconv"
	//"time"
	"encoding/json"
	//"math/rand"
	"net/http"

	"github.com/gorilla/mux"

	"../common"
)


func root(w http.ResponseWriter, r *http.Request) {

	if (r.ContentLength < 6) {
		w.WriteHeader(400)
		w.Write([]byte("400 bad request"))
	} else {
		w.Write([]byte("hello root\n"))
	}
}


func all(w http.ResponseWriter, r *http.Request) {
	var res string
	var tot int
	var err error
	var ja map[string]interface{}


	fmt.Printf("req: %s\n", r.URL)

	w.Header().Set("Content-Type", "text/json")


	res, _ = common.Xhttp("GET", common.I_PORT, "/all", "")
	//fmt.Printf("res: %+v\n", res)

	err = json.Unmarshal([]byte(res), &ja)
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "panic: bad json from micro [%s]\n", err.Error())

		w.WriteHeader(500)
		w.Write([]byte("{\"status\": 500, \"msg\": \"Inernal server error\"}"))

		return
	}
	if (int(ja["status"].(float64)) != 200) {
		fmt.Fprintf(os.Stderr, "panic: micro error %d [%s]\n", int(ja["status"].(float64)), ja["msg"])

		w.WriteHeader(500)
		w.Write([]byte("{\"status\": 500, \"msg\": \"Inernal server error\"}"))

		return
	}
	//fmt.Printf("%+v\n", ja)

	tot = int(ja["total"].(float64))

	jm, err := json.Marshal(ja["results"])

	b := fmt.Sprintf("{\"status\": 200, \"msg\": \"ok\", \"total\": %d, \"results\": %s}", tot, jm)
	w.Write([]byte(b))

}


func client(w http.ResponseWriter, r *http.Request) {
}


func last(w http.ResponseWriter, r *http.Request) {
}


func today(w http.ResponseWriter, r *http.Request) {
}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", root)
	router.HandleFunc("/all", all)
	router.HandleFunc("/client", client)
	router.HandleFunc("/client/", client)
	router.HandleFunc("/client/{client}", client)
	router.HandleFunc("/last", last)
	router.HandleFunc("/last/", last)
	router.HandleFunc("/last/{sec}", last)
	router.HandleFunc("/today", today)

	b := fmt.Sprintf(":%d", common.R_PORT)
	http.ListenAndServe(b, router)

}
