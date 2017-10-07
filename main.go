package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"os"
	// "log"
	"encoding/json"
	"net/http"
	"strings"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Room struct {
    Id   string
	Room_type string
}

type RoomInfo struct {
    Id   string
	start_date string
	end_date string
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("udeain"))
}
func splitDate(date string)(retDate map[string]string){
	dateArray := strings.Split(date, "-")
	//implementar validaciones para saber si existen los 3 atributos de la fecha
	retDate = map[string]string{
		"day":dateArray[0],
		"month":dateArray[1],
		"year":dateArray[2],
	}
	return
}
func getRooms(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	//split the arrive date to get all the info
	arriveDate := vars["arriveDate"]
	arriveDateObj := splitDate(arriveDate)

	//split the arrive date to get all the info
	leaveDate := vars["leaveDate"]
	leaveDateObj := splitDate(leaveDate)

	city := vars["city"]
	hosts := vars["hosts"]
	roomType := vars["roomType"]

	println("searching.--.....----.")
	println("arriveDate",arriveDateObj["year"],arriveDateObj["month"],arriveDateObj["day"])
	println("leaveDate",leaveDateObj["year"],leaveDateObj["month"],leaveDateObj["day"])
	println("city",city)
	println("hosts",hosts)
	println("roomType",roomType)
	

	session, err := mgo.Dial("mongodb://udeain:udeainmongodb@ds157444.mlab.com:57444/heroku_4r2js6cs")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("heroku_4r2js6cs").C("rooms")

	// result := Room{}
	var roomsObj []bson.M
	err = c.Find(bson.M{"room_type": roomType, "city":city, "available":true}).All(&roomsObj)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("not found"))
		return
	}
	respuesta, err :=  json.Marshal(roomsObj)
	if err != nil {
		w.WriteHeader(405)
		w.Write([]byte("unable to get room"))
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(respuesta)
}

func getRoomsAvailable(w http.ResponseWriter, r *http.Request){

	city := "05001"
	roomType := "s"

	// establecer conexión
	session, err := mgo.Dial("mongodb://udeain:udeainmongodb@ds157444.mlab.com:57444/heroku_4r2js6cs")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	c := session.DB("heroku_4r2js6cs").C("rooms_info")

	w.Header().Set("Content-Type", "application/json")

	// filtro de resultados
	var rooms []bson.M
	err = c.Find(bson.M{"room_type": roomType, "city":city, "available":true}).All(&rooms)

	respuesta, err :=  json.Marshal(rooms)
	if err != nil {
		w.WriteHeader(405)
		w.Write([]byte("unable to get room"))
		return
	}

	var rooms_data []Room
	json.Unmarshal(respuesta, &rooms_data)

	for item := range rooms_data{
		w.Write( []byte("Room")) 
		w.Write( []byte(rooms_data[item].Room_type) )
	}
	

}

//mongodb://udeain:udeainmongodb@ds157444.mlab.com:57444/heroku_4r2js6cs
//http://localhost:8080/api/v1/rooms/arrive_date/01-01-2017/leave_date/02-02-2017/city/05001/hosts/3/room_type/l
func main(){
	fmt.Println("start server 8080")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/api/v1/rooms/arrive_date/{arriveDate}/leave_date/{leaveDate}/city/{city}/hosts/{hosts}/room_type/{roomType}", getRooms).Methods("GET")
	r.HandleFunc("/api/v1/rooms_info", getRoomsAvailable).Methods("GET")
	
	http.Handle("/", r)
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe("0.0.0.0:"+port, nil)
}