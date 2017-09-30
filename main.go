package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"os"
	// "log"
	// "encoding/json"
	"net/http"
	// "io/ioutil"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

///Users/usuario/go/src/github.com/gorilla/mux/
func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func getRooms(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	arriveDate := vars["arriveDate"]
	leaveDate := vars["leaveDate"]
	city := vars["city"]
	hosts := vars["hosts"]
	roomType := vars["roomType"]
	println("arriveDate",arriveDate)
	println("leaveDate",leaveDate)
	println("city",city)
	println("hosts",hosts)
	println("roomType",roomType)
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

//http://localhost:8080/api/v1/rooms/arrive_date/01-01-2017/leave_date/02-02-2017/city/05001/hosts/3/room_type/l
func main(){
	fmt.Println("start server 8080")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/api/v1/rooms/arrive_date/{arriveDate}/leave_date/{leaveDate}/city/{city}/hosts/{hosts}/room_type/{roomType}", getRooms).Methods("GET")
	
	http.Handle("/", r)
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
