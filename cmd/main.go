package main

import (
	"DataMapper/internal/api"
)

func main() {
	request := api.Request{Url: "http://localhost:8080/get", Params: map[string]string{"id": "12345"}, Method: api.GET, Body: ""}
	request.Do()
}
