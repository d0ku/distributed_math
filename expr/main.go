//go:generate ragel -Z lex.rl
//go:generate goyacc -o parser.go -p "pars" parser.y
//#go:generate stringer -type=CommandType

package main

import (
	"fmt"
)

/*
func expressionHandler(w http.ResponseWriter, r *http.Request, e *base.ExpressionSingle) {
	var result int
	fmt.Println(e)

	// DO MATCHING HERE
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
*/

func main() {
	//http.HandleFunc("/", base.ExprWrapper(expressionHandler))
	//http.ListenAndServe(":8080", nil)
	content := []byte("  2 + 2 + 2 * (2+3)")
	parsParse(newLexer(content))
	fmt.Println(result)
}
