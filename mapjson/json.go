package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	soi := map[string]interface{}{
		"id":   1,
		"name": "My",
		"age":  18,
	}

	//chuyen doi du lieu
	data, _ := json.Marshal(soi)
	var myJsonTring = string(data)
	fmt.Println(myJsonTring)
}
