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

func TestJsonArrayComplext(t *testing.T) {
	customer := []Customer{
		{Name: "Adit",
			Age:       26,
			Addresses: []Address{{Village: "Petahunan RT 01/ RW 01", District: "Sempor", Regency: "Kebumen", Province: "Jawa Tengah"}},
			Hobbies:   []string{"Gaming", "Coding", "Touring"}},
		{Name: "Hakim",
			Age:       26,
			Addresses: []Address{{Village: "Gunung pati", District: "Semarang barat", Regency: "Semarang", Province: "Jawa Tengah"}},
			Hobbies:   []string{"Gaming", "Coding", "Touring"}},
		{Name: "Dio",
			Age:       26,
			Addresses: []Address{{Village: "Banyu Biru", District: "Ambarawa", Regency: "Kab. Semarang", Province: "Jawa Tengah"}},
			Hobbies:   []string{"Gaming", "Coding", "Touring"}},
		{Name: "Danang",
			Age:       26,
			Addresses: []Address{{Village: "Purwodadi", District: "Purwodadi", Regency: "Purwodadi", Province: "Jawa Tengah"}},
			Hobbies:   []string{"Gaming", "Coding", "Touring"}},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestDecodeArrayJsonComplex(t *testing.T) {
	jsonString := `[{"nama":"Adit","age":26,"addresses":[{"village":"Petahunan RT 01/ RW 01","district":"Sempor","regency":"Kebumen","province":"Jawa Tengah"}],"hobbies":["Gaming","Coding","Touring"]},{"nama":"Hakim","age":26,"addresses":[{"village":"Gunung pati","district":"Semarang barat","regency":"Semarang","province":"Jawa Tengah"}],"hobbies":["Gaming","Coding","Touring"]},{"nama":"Dio","age":26,"addresses":[{"village":"Banyu Biru","district":"Ambarawa","regency":"Kab. Semarang","province":"Jawa Tengah"}],"hobbies":["Gaming","Coding","Touring"]},{"nama":"Danang","age":26,"addresses":[{"village":"Purwodadi","district":"Purwodadi","regency":"Purwodadi","province":"Jawa Tengah"}],"hobbies":["Gaming","Coding","Touring"]}]`
	jsonByte := []byte(jsonString)

	customer := &[]Customer{}

	json.Unmarshal(jsonByte, customer)
	fmt.Println(customer)

}
