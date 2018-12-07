package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"encoding/json"
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
	var client int
	var res string
	var tot int
	var err error
	var ja map[string]interface{}
	var jr []interface{}


	fmt.Printf("req: %s\n", r.URL)

	w.Header().Set("Content-Type", "text/json")

	vars := mux.Vars(r)

	client, _ = strconv.Atoi(vars["client"])

	if (client == 0) {
		w.Write([]byte("{\"status\": 200, \"msg\": \"id not found\"}"))
		return
	}


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

	tot = 0

	for _, v := range(ja["results"].([]interface{})) {

		if (int(v.(map[string]interface{})["client"].(float64)) != client) {
			continue
		}

		jr = append(jr, v)
		tot += 1

	}
	//fmt.Printf("%v\n", jr)

	jm, err := json.Marshal(jr)

	b := fmt.Sprintf("{\"status\": 200, \"msg\": \"ok\", \"total\": %d, \"results\": %s}", tot, jm)
	w.Write([]byte(b))

}


func last(w http.ResponseWriter, r *http.Request) {
	var sec int
	var res string
	var tot int
	var err error
	var ja map[string]interface{}
	var jr []interface{}


	fmt.Printf("req: %s\n", r.URL)

	w.Header().Set("Content-Type", "text/json")

	vars := mux.Vars(r)

	sec, _ = strconv.Atoi(vars["sec"])

	if (sec <= 0) {
		w.Write([]byte("{\"status\": 400, \"msg\": \"bad request\"}"))
		return
	}


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

	tot = 0
	ti := int(time.Now().Unix()) - sec

	for _, v := range(ja["results"].([]interface{})) {

		if (ti > int(v.(map[string]interface{})["utime"].(float64))) {
			continue
		}

		jr = append(jr, v)
		tot += 1

	}
	//fmt.Printf("%v\n", jr)

	jm, err := json.Marshal(jr)

	b := fmt.Sprintf("{\"status\": 200, \"msg\": \"ok\", \"total\": %d, \"results\": %s}", tot, jm)
	w.Write([]byte(b))

}


func today(w http.ResponseWriter, r *http.Request) {
	var res string
	var tot int
	var err error
	var ja map[string]interface{}
	var jr []interface{}


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

	tot = 0

	now := time.Now()
	ty, tm, td := now.Date()
	ti := time.Date(ty, tm, td, 0, 0, 0, 0, time.Local).Unix()
	fmt.Printf("ti %d\n", ti)

	for _, v := range(ja["results"].([]interface{})) {

		if (ti > int64(v.(map[string]interface{})["utime"].(float64))) {
			continue
		}

		jr = append(jr, v)
		tot += 1

	}
	//fmt.Printf("%v\n", jr)

	jm, err := json.Marshal(jr)

	b := fmt.Sprintf("{\"status\": 200, \"msg\": \"ok\", \"total\": %d, \"results\": %s}", tot, jm)
	w.Write([]byte(b))

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
