package main

import (
	"fmt"
	"log"
	"net/http"
)

func test(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println(req.Body)
	fmt.Println("Hello-world")

	fmt.Println(req.FormValue("param1"))
	//LOG: map[{"test": "that"}:[]]
	// var t test_struct
	// for key, _ := range req.Form {
	// 	log.Println(key)
	// 	//LOG: {"test": "that"}
	// 	err := json.Unmarshal([]byte(key), &t)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 	}
	// }
	// log.Println(t.Test)
	//LOG: that
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
