package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

var data = ` {
	"name": "arian",
	"type": "deposit",
	"amount": 1000.56
	}
`

type Data struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	reader := bytes.NewBufferString(data)

	decoder := json.NewDecoder(reader)

	request := &Data{} // req for amount

	err := decoder.Decode(request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", request)

	// create response
	var prevBalance float64 = 8500000
	response := map[string]interface{}{
		"ok":      true,
		"code":    200,
		"balance": prevBalance + request.Amount,
	}
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(response); err != nil {
		panic(err)
	}
}
