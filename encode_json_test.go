package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	LogJson("Gita")
	LogJson(1)
	LogJson(0.5)
	LogJson(true)
	LogJson([]string{"Bakso", "Mie Ayam", "Geprek"})

}
