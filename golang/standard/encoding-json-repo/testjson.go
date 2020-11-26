package testjson

import (
	"encoding/json"
	"fmt"
	"strings"
)

func TestJson() {
	var t_result map[string]interface{}

	t_result = map[string]interface{}{
		"name": "andrew",
		"age":  18,
	}

	d1, _ := json.Marshal(&t_result)
	fmt.Println(d1)

	var result map[string]interface{}
	d := json.NewDecoder(strings.NewReader(string(d1)))
	d.UseNumber()
	err := d.Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%T\n", result["age"])
}
