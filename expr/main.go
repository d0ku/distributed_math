//go:generate ragel -Z lex.rl
//go:generate goyacc -o parser.go -p "pars" parser.y
//#go:generate stringer -type=CommandType

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/d0ku/distributed_math/base"
)

func expressionHandler(w http.ResponseWriter, r *http.Request, e *base.ExpressionSingle) {
	var jsn []byte
	var err error

	log.Println("Got request: " + e.Content)
	content := []byte(e.Content)
	if parsParse(newLexer(content)) == 0 {
		res := base.Response{base.Success, result}
		jsn, err = json.Marshal(res)
		if err != nil {
			base.CreateError(w, r)
		}
		fmt.Fprintf(w, string(jsn))
	} else {
		res := base.Response{base.FailureInterpretExpression, result}
		jsn, err = json.Marshal(res)
		if err != nil {
			base.CreateError(w, r)
		}
		fmt.Fprintf(w, string(jsn))
	}
	log.Println("Send response:" + string(jsn))
}

func main() {
	http.HandleFunc("/", base.ExprWrapper(expressionHandler))
	http.ListenAndServe(":8080", nil)
}
