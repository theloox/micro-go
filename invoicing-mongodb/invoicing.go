package main

import (
    "fmt"
	"os"
	"time"
    "encoding/json"
	"math/rand"
    "net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

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

func create(w http.ResponseWriter, r *http.Request) {

	var b []byte
	var in map[string]interface{}
	//var rec common.Invoice
	var rid bson.M
	var id int


	fmt.Printf("req: %s\n", r.URL)

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

	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"val": 1}},
		ReturnNew: true,
	}

	_, err = db.Find(bson.M{"_id": "counter"}).Apply(change, &rid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "panic: Can't insert in db %s\n", err)

		w.WriteHeader(500)
		w.Write([]byte("{\"status\": 500, \"msg:\": \"Inernal server error\"}"))

		return
	}

	id = rid["val"].(int)

	tm := time.Now()

	//rec = common.Invoice{_id: id, Id: id, Uid: rand.Uint32(), Utime: tm.Unix(), Client: in["client"].(float64), Items: in["items"].(float64), Amount: in["amount"].(float64)}

	m := bson.M{
		"_id": id,
		"id": id,
		"uid": rand.Uint32(),
		"utime": tm.Unix(),
		"client": in["client"].(float64),
		"items": in["items"].(float64),
		"amount": in["amount"].(float64),
	}
	//fmt.Printf("m: %v\n\n", m);

	err = db.Insert(&m)
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

	fmt.Printf("req: %s\n", r.URL)

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
	session, err := mgo.Dial("localhost/invoicing")
	if err != nil {
		fmt.Fprintf(os.Stderr, "panic: db connection failed %s\n", err)
		os.Exit(2)
	}

	var res map[string]string
	db := session.DB("invoicing").C("invoices")
	db.FindId("counter").One(&res)

	if (res["_id"] == "") {
		fmt.Printf("info: init DB\n\n")

		err = db.Insert(bson.M{"_id": "counter", "val": 0})

		if err != nil {
			fmt.Fprintf(os.Stderr, "panic: db counter failed %s\n", err)
			os.Exit(2)
		}
	}

	session.Close()

    http.ListenAndServe(":8080", router)
}
