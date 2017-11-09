package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	/* "net/http"
	"strings"
	"io/ioutil"	*/)

func Test1(t *testing.T) {
	expected := "udeain"
	actual := "udeain"
	if actual != expected {
		t.Error("Test failed")
	}
}

func Test2(t *testing.T) {

	// formato de json enviado por medio de formulario web
	jsonDatos := []byte(`{"arrive_date":"2017-10-25","leave_date":"2017-10-26","room_type":"s","capacity":1,"beds":{"simple":1,"double":0},"hotel_id":"udeain_medellin",
		"user":{"doc_type":"CC","doc_id":"11521777","email":"cjmo@gmail.com","phone_number":"4448787"}}`)

	// Procesar datos recibidos
	var raw map[string]interface{}
	json.Unmarshal(jsonDatos, &raw)

	salida, _ := json.Marshal(raw["arrive_date"])

	expected := string("2017-10-25")

	arrive_date := string(salida)
	arrive_date = strings.Replace(arrive_date, "\"", "", -1)

	//fmt.Println(arrive_date)

	// prueba de verificación de fecha de llegada
	if arrive_date != expected {
		t.Error("Test failed para dato arrive_date")
	}

	// prueba de verificación de fechas (que la de llegada sea inferior a la de salida)
	salida, _ = json.Marshal(raw["leave_date"])
	leave_date := string(salida)
	leave_date = strings.Replace(leave_date, "\"", "", -1)

	if arrive_date > leave_date {
		t.Error("Test failed: fecha de llegada es posterior a la de salida")
	}

	// prueba de tipo de habitación
	salida, _ = json.Marshal(raw["room_type"])
	room_type := string(salida)
	room_type = strings.Replace(room_type, "\"", "", -1)
	room_type = strings.ToUpper(room_type)
	if room_type != "S" && room_type != "L" {
		t.Error("Test failed: Tipo de habitación distinta a las soportadas: S y L")
	}

	// prueba de capacidad de habitación (personas)
	salida, _ = json.Marshal(raw["capacity"])
	capacity := string(salida)
	capacity = strings.Replace(capacity, "\"", "", -1)
	capacity_number, err := strconv.Atoi(capacity)
	//fmt.Println(err)
	if capacity_number < 1 || err != nil {
		t.Error("Test failed: La capacidad de la habitación debe ser mínimo de 1 persona")
	}

}

func Test3(t *testing.T) {
	// prueba de puerto de comunicación
	puerto := os.Getenv("PORT")
	fmt.Println("Puerto: " + puerto)

	puerto_esperado := "8080"
	if puerto != puerto_esperado {
		t.Error("Test failed: Puerto de comunicación del servidor diferente al esperado")
	}
}
