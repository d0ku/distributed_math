package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/d0ku/distributed_math/base"
)

func expressionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsn, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(jsn))
	if err != nil {
		r := base.Response{base.FailureReadBody, 1}
		m, err := json.Marshal(r)
		if err != nil {
			log.Println("Could not read body json")
		}
		fmt.Fprintf(w, string(m))
		return
	}

	req := base.Expression{}
	err = json.Unmarshal(jsn, &req)
	if err != nil {
		r := base.Response{base.FailureParseJSON, 2}
		m, err := json.Marshal(r)
		if err != nil {
			log.Println("Could not parse response json")
		}
		fmt.Fprintf(w, string(m))
		return
	}

	res := base.Response{base.Success, req.Solve()}
	jsn, err = json.Marshal(res)
	if err != nil {
		r := base.Response{base.FailureCreateJSON, 3}
		m, err := json.Marshal(r)
		if err != nil {
			log.Println("Could not create response json")
		}
		fmt.Fprintf(w, string(m))
		return
	}

	fmt.Fprintf(w, string(jsn))
}

func main() {
	exp := base.Expression{1, 2, '*'}
	jsn, _ := json.Marshal(exp)
	fmt.Println(string(jsn))
	http.HandleFunc("/", expressionHandler)
	http.ListenAndServe(":8080", nil)
}
