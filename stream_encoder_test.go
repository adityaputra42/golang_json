package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		Name: "Putra",
		Age:  32,
		Addresses: []Address{
			{Village: "Sambiroto", District: "Tembalang", Regency: "Kota Semarang", Province: "Jawa Tengah"},
		},
		Hobbies: []string{"Futsal", "Coding"},
	}
	encoder.Encode(customer)
	

}
