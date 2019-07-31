package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type State int
type Operator int

const (
	Success State = iota
	FailureReadBody
	FailureParseJSON
	FailureCreateJSON

	Add Operator = iota
	Sub
	Mult
	Div
)

type Response struct {
	Status State
	Result int
}

type ArgumentNumber struct {
	Value int
}

func (a *ArgumentNumber) IsNumber() bool {
	return true
}

type ArgumentExpr struct {
	Expr ExpressionSingle
}

func (a *ArgumentExpr) IsNumber() bool {
	return false
}

type Argument interface {
	// IsNumber allows to differentiate between number (int value) and not yet resolved expression.
	IsNumber() bool
}

type ExpressionSingle struct {
	Content string
}

func (e *ExpressionSingle) HasOperator() bool {
	return false
}

type ExpressionOperation struct {
	First  Argument
	Second Argument
	Op     Operator
}

func (e *ExpressionOperation) HasOperator() bool {
	return true
}

type Expression interface {
	HasOperator() bool
}

func ParseError(w http.ResponseWriter, r *http.Request) {
	log.Println("Could not parse request json")
	res := Response{FailureParseJSON, 0}
	m, err := json.Marshal(res)
	if err != nil {
		log.Println("Could not create error json")
	}
	fmt.Fprintf(w, string(m))
}

func CreateError(w http.ResponseWriter, r *http.Request) {
	log.Println("Could not create response json")
	res := Response{FailureCreateJSON, 0}
	m, err := json.Marshal(res)
	if err != nil {
		log.Println("Could not create error json")
	}
	fmt.Fprintf(w, string(m))
}

func HandlerWrapper(f func(w http.ResponseWriter, r *http.Request, req Expression)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jsn, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(jsn))
		if err != nil {
			log.Println("Could not read body json")
			r := Response{FailureReadBody, 0}
			m, err := json.Marshal(r)
			if err != nil {
				log.Println("Could not create error json")
			}
			fmt.Fprintf(w, string(m))
			return
		}

		var req Expression
		reqOne := ExpressionOperation{}
		err = json.Unmarshal(jsn, &reqOne)
		if err != nil {
			reqTwo := ExpressionSingle{}
			err = json.Unmarshal(jsn, &reqTwo)
			if err != nil {
				ParseError(w, r)
				return
			}
			req = &reqTwo
		} else {
			req = &reqOne
		}

		f(w, r, req)
	}
}

func SolveExpression(url string, exp Expression) int {
	jsonStr, err := json.Marshal(exp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	res := Response{}
	json.Unmarshal(body, &res)
	return res.Result
}
