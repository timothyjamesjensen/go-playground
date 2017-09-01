package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Response struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	City string `json:"city"`
	State string `json:"state"`
}

func main() {
	resp, err := http.Get("")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()	

	r := new(Response)

	json.NewDecoder(resp.Body).Decode(&r)

	fmt.Println(r)
}
