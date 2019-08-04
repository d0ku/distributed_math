package main

import (
	"fmt"
	"os"

	"github.com/d0ku/distributed_math/base"
)

func main() {
	url := os.Args[1]
	fmt.Println("URL:>", url)

	first := base.ArgumentUnion{IsNumber: true, Number: 1}
	second := base.ArgumentUnion{IsNumber: true, Number: 2}
	exp := &base.ExpressionOperation{first, second, base.Add}

	base.SolveExpression(url, exp)
}
