package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}

func raw() ([]Data, error) {
	r, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	// Paylaod
	var postPayload Payload
	err = json.Unmarshal(r, &postPayload.Data)
	if err != nil {
		return nil, err
	}

	// Print the payload data
	fmt.Println("Payload data:", postPayload.Data)

	return postPayload.Data, nil
}

func GetAll() ([]Data, error) {
	res, err := raw()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetById(id int) (any, error) {
	res, err := raw()
	if err != nil {
		return nil, err
	}
	if id > len(res) {
		res := make([]string, 0)
		return res, nil
	}
	return res[id], nil
}
