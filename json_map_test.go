package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJsonDecode(t *testing.T) {
	jsonString := `[{"nama":"Adit","age":26,"addresses":[{"village":"Petahunan RT 01/ RW 01","district":"Sempor","regency":"Kebumen","province":"Jawa Tengah"}],"hobbies":["Gaming","Coding","Touring"]},{"nama":"Hakim","age":26,"addresses":[{"village":"Gunung pati","district":"Semarang barat","regency":"Semarang","province":"Jawa Tengah"}],"hobbies":["Gaming","Coding","Touring"]},{"nama":"Dio","age":26,"addresses":[{"village":"Banyu Biru","district":"Ambarawa","regency":"Kab. Semarang","province":"Jawa Tengah"}],"hobbies":["Gaming","Coding","Touring"]},{"nama":"Danang","age":26,"addresses":[{"village":"Purwodadi","district":"Purwodadi","regency":"Purwodadi","province":"Jawa Tengah"}],"hobbies":["Gaming","Coding","Touring"]}]`
	jsonByte := []byte(jsonString)

	var result []map[string]interface{}
	json.Unmarshal(jsonByte, &result)
	for _, value := range result {
		fmt.Println(value["nama"])
		fmt.Println(value["age"])
	}
}

func TestMapJsonEncode(t *testing.T) {
	product := []map[string]interface{}{
		{"id": 1, "name": "Air Jordan", "price": 1750000},
		{"id": 2, "name": "Nike", "price": 370000},
		{"id": 3, "name": "Adidas", "price": 275000},
		{"id": 4, "name": "Cheetah", "price": 450000},
		{"id": 5, "name": "Ortuse Eight", "price": 150000},
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))

}
