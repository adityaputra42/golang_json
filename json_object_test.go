package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Village  string `json:"village"`
	District string `json:"district"`
	Regency  string `json:"regency"`
	Province string `json:"province"`
}

type Customer struct {
	Name      string    `json:"nama"`
	Age       int16     `json:"age"`
	Addresses []Address `json:"addresses"`
	Hobbies   []string  `json:"hobbies"`
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		Name:      "Aditya",
		Age:       27,
		Addresses: []Address{Address{Village: "Petahunan RT 01/ RW 01", District: "Sempor", Regency: "Kebumen", Province: "Jawa Tengah"}},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}
