package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"name":"Aditya","age":27,"address":{"village":"Petahunan RT 01/ RW 01","district":"Sempor","regency":"Kebumen","province":"Jawa Tengah"}}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}

	json.Unmarshal(jsonByte, customer)
	fmt.Println(customer)
}
