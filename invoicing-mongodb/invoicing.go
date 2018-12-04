package main

import (
    "fmt"
	"os"
	"strconv"
	"time"
    "encoding/json"
	"math/rand"
    "net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"../common"
)

type Element struct {
	ID int		`json:"id,omitempty`
	Times int	`json:"times,omitempty`
	Name string	`json:"name,omitempty`
}

type Elements []Element

type Message struct {
	Msg string
}

var list = Elements {
	Element{ID: 1, Times: 100, Name: "this is 1"},
	Element{ID: 2, Times: 2000, Name: "maybe 2"},
	Element{ID: 3, Times: 30000, Name: "and 3"},
}


func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><body>")

    fmt.Fprintf(w, "Hello my baby Hello my honey, %s", r.URL.Path[1:])
    fmt.Fprintf(w, "<br>", )
    fmt.Fprintf(w, "now go to <a href=\"/app\">/app</a>")

	fmt.Fprintf(w, "</body></html>")

}

func app(w http.ResponseWriter, r *http.Request) {

    b, err := json.Marshal(list)

    if err != nil {
        panic(err)
    }

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(b)
}

func appid(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
    id_ := vars["appid"]

	id, _ := strconv.Atoi(id_)

	var ret Element

	for _, e := range list {
		if e.ID == id {
			ret = e
		}
	}

    b, err := json.Marshal(ret)

    if err != nil {
        panic(err)
    }

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(b)
}

func about(w http.ResponseWriter, r *http.Request) {

    m := Message{"webrest, build v0.0"}
    b, err := json.Marshal(m)

    if err != nil {
        panic(err)
    }

     w.Write(b)
}


func root(w http.ResponseWriter, r *http.Request) {

	if (r.ContentLength < 6) {
		w.WriteHeader(400)
		w.Write([]byte("400 bad request"))
	} else {
		w.Write([]byte("hello root\n"))
	}
}

func create(w http.ResponseWriter, r *http.Request) {

	var b []byte
	var in map[string]interface{}


	w.Header().Set("Content-Type", "text/json")

	if (r.Body == nil) {
		w.WriteHeader(400)
		w.Write([]byte("{\"status\": 400, \"msg:\": \"bad request\""))
		return
	}

	// db
	session, err := mgo.Dial("localhost/invoicing")
	if err != nil {
		fmt.Fprintf(os.Stderr, "panic: db connection failed %s\n", err)

		w.WriteHeader(500)
		w.Write([]byte("{\"status\": 500, \"msg:\": \"Inernal server error\"}"))

		return
	}
	defer session.Close()

	db := session.DB("invoicing").C("invoices")

	d := json.NewDecoder(r.Body)

	err = d.Decode(&in)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"status\": 400, \"msg:\": \"bad request\"}"))
		//w.Write([]byte(err.Error()))
		return
	}
	r.Body.Close()

	var rec common.Invoice

	tm := time.Now()
	rec = common.Invoice{Uid: rand.Uint32(), Utime: tm.Unix(), Client: in["client"].(float64), Items: in["items"].(float64), Amount: in["amount"].(float64)}

	err = db.Insert(&rec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "panic: Can't insert in db\n")

		w.WriteHeader(500)
		w.Write([]byte("{\"status\": 500, \"msg:\": \"Inernal server error\"}"))

		return
	}

	w.Write([]byte("{\"status\": 200, \"msg\": \"ok\"}"))

	//b = []byte(fmt.Sprintf("\n%+v\n%v\n", in, rec))
	w.Write(b)

}

func rd(w http.ResponseWriter, r *http.Request) {

}

func update(w http.ResponseWriter, r *http.Request){

}

func del(w http.ResponseWriter, r *http.Request) {

}

func all(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/json")

	// db
	session, err := mgo.Dial("localhost/invoicing")
	if err != nil {
		fmt.Fprintf(os.Stderr, "panic: db connection failed %s\n", err)

		w.WriteHeader(500)
		w.Write([]byte("{\"status\": 500, \"msg:\": \"Inernal server error\"}"))

		return
	}
	defer session.Close()

	db := session.DB("invoicing").C("invoices")


	var res []common.Invoice
	err = db.Find(nil).All(&res)

	jn, _ := json.Marshal(res)

	w.Write([]byte("{\"status\": 200, \"msg:\": \"ok\", \"results\": "))
	w.Write(jn)
	w.Write([]byte("}"))
}


func main() {
	router := mux.NewRouter()

    router.HandleFunc("/", root)
    router.HandleFunc("/create", create)
    router.HandleFunc("/read/{appid}", rd)
    router.HandleFunc("/update", update)
    router.HandleFunc("/delete", del)
    router.HandleFunc("/all", all)

	// db
	_, err := mgo.Dial("localhost/invoicing")
	if err != nil {
		fmt.Fprintf(os.Stderr, "panic: db connection failed %s\n", err)
		os.Exit(2)
	}

    http.ListenAndServe(":8080", router)
}
