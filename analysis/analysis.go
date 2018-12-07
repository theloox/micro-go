package main

import (
	"fmt"
	"os"
	"strconv"
	//"time"
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


func abnormal(w http.ResponseWriter, r *http.Request) {

}

func duplicates(w http.ResponseWriter, r *http.Request) {

}

func same(w http.ResponseWriter, r *http.Request) {

}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", root)
	router.HandleFunc("/abnormal", abnormal)
	router.HandleFunc("/duplicates", duplicates)
	router.HandleFunc("/same", same)

	b := fmt.Sprintf(":%d", common.R_PORT)
	http.ListenAndServe(b, router)

}
