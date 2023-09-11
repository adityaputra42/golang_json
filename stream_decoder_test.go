package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {

	file, _ := os.Open("customer.json")

	decoder := json.NewDecoder(file)

	customers := &[]Customer{}
	decoder.Decode(customers)
	for _, value := range *customers {
		fmt.Println(value.Name)
	}

}
