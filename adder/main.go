package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/d0ku/distributed_math/base"
)

func expressionHandler(w http.ResponseWriter, r *http.Request, req *base.ExpressionOperation) {
	log.Println("Adder got request")
	fChan := make(chan int, 1)
	sChan := make(chan int, 1)
	if req.First.IsNumber {
		fChan <- req.First.Number
	} else {
		sChan <- base.SolveExpression("http://localhost:8000", &req.First.Expr)
	}

	if req.Second.IsNumber {
		sChan <- req.Second.Number
	} else {
		sChan <- base.SolveExpression("http://localhost:8000", &req.Second.Expr)
	}

	res := base.Response{base.Success, <-fChan + <-sChan}
	log.Println("Adder solved task")
	jsn, err := json.Marshal(res)
	if err != nil {
		base.CreateError(w, r)
	}

	fmt.Fprintf(w, string(jsn))
}

func main() {
	http.HandleFunc("/", base.OperationWrapper(expressionHandler))
	http.ListenAndServe(":8081", nil)
}
