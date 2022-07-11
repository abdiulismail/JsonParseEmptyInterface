package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var f interface{}
	_ = json.Unmarshal(body, &f)
	fmt.Printf("%+v",f)

}

