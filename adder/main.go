package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/d0ku/distributed_math/base"
)

func expressionHandler(w http.ResponseWriter, r *http.Request, reqRec base.Expression) {
	log.Println("Adder got request")
	req := reqRec.(*base.ExpressionOperation)
	fChan := make(chan int, 1)
	sChan := make(chan int, 1)
	if req.First.IsNumber() {
		fChan <- (req.First.(*base.ArgumentNumber)).Value
	} else {
		sChan <- base.SolveExpression("http://localhost:8000", &(req.First.(*base.ArgumentExpr)).Expr)
	}

	if req.Second.IsNumber() {
		sChan <- (req.Second.(*base.ArgumentNumber)).Value
	} else {
		sChan <- base.SolveExpression("http://localhost:8000", &(req.Second.(*base.ArgumentExpr)).Expr)
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
	http.HandleFunc("/", base.HandlerWrapper(expressionHandler))
	http.ListenAndServe(":8081", nil)
}
