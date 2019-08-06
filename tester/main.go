package main

import (
	"fmt"
	"os"

	"github.com/d0ku/distributed_math/base"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: app_name url expression")
		fmt.Println("eg ./tester https://math.d0ku.org '2+2*2'")
		return
	}
	url := os.Args[1]
	expr := os.Args[2]
	fmt.Println("URL:>", url)

	exp := &base.ExpressionSingle{expr}

	base.SolveExpression(url, exp)
}
