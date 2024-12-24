package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var object interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	json.Unmarshal(b, &object)
	fmt.Println(object)
}
