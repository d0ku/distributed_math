package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/d0ku/distributed_math/base"
)

func expressionHandler(w http.ResponseWriter, r *http.Request, e base.Expression) {
	var result int
	if e.HasOperator() {
		exp := e.(*base.ExpressionOperation)
		switch exp.Op {
		case base.Add:
			result = base.SolveExpression("http://localhost:8081", exp)
			// ask adder
		}
		// assign result =
		// send result in response
	} else {
		// special case
		exp := e.(*base.ExpressionSingle)
		res, err := strconv.ParseInt(exp.Content, 10, 64)
		if err != nil {
			panic(err)
		}
		result = int(res)
	}
	res := base.Response{base.Success, result}
	jsn, err := json.Marshal(res)
	if err != nil {
		base.CreateError(w, r)
	}

	fmt.Fprintf(w, string(jsn))
}

func main() {
	http.HandleFunc("/", base.HandlerWrapper(expressionHandler))
	http.ListenAndServe(":8080", nil)
}
