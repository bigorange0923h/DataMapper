package main

import (
	"DataMapper/internal/api"
)

func main() {
	request := api.Request{Url: "http://localhost:8080/post", Params: map[string]string{"id": "12345"}, Method: api.POST, Body: "{\"test\":\"123\"}"}
	request.Do()
}
