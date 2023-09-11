package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		Name:      "Gita",
		Age:       26,
		Addresses: []Address{{Village: "Petahunan RT 01/ RW 01", District: "Sempor", Regency: "Kebumen", Province: "Jawa Tengah"}, {}, {}, {}},
		Hobbies:   []string{"Gaming", "Coding", "Touring"},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestDecodeArrayJson(t *testing.T) {
	jsonString := `{"nama":"Gita","age":26,"addresses":[{"village":"Petahunan RT 01/ RW 01","district":"Sempor","regency":"Kebumen","province":"Jawa Tengah"},{"village":"","district":"","regency":"","province":""},{"village":"","district":"","regency":"","province":""},{"village":"","district":"","regency":"","province":""}],"hobbies":["Gaming","Coding","Touring"]}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}

	json.Unmarshal(jsonByte, customer)
	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Addresses)
	fmt.Println(customer.Hobbies)
}
