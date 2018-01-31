package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"./utils"
)

type test_struct struct {
	Test  string
	Test1 string
}

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t test_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	utils.INFO.Println(t.Test)
	utils.TRACE.Println(t.Test1)
	utils.WARNING.Println(t.Test1)
	utils.ERROR.Println(t.Test)
}

func main() {
	utils.Init()
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
