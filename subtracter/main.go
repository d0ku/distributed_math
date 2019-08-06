package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/d0ku/distributed_math/base"
)

func expressionHandler(w http.ResponseWriter, r *http.Request, req *base.ExpressionOperation) {
	log.Println("Subtracter got request")
	fChan := make(chan int, 1)
	sChan := make(chan int, 1)
	if req.First.IsNumber {
		fChan <- req.First.Number
	} else {
		sChan <- base.SolveExpression(os.Getenv("EXPR_URL"), &req.First.Expr)
	}

	if req.Second.IsNumber {
		sChan <- req.Second.Number
	} else {
		sChan <- base.SolveExpression(os.Getenv("EXPR_URL"), &req.Second.Expr)
	}

	res := base.Response{base.Success, <-fChan - <-sChan}
	log.Println("Subtracter solved task")
	jsn, err := json.Marshal(res)
	if err != nil {
		base.CreateError(w, r)
	}

	fmt.Fprintf(w, string(jsn))
}

func main() {
	http.HandleFunc("/", base.OperationWrapper(expressionHandler))
	http.ListenAndServe(":8082", nil)
}
