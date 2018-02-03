package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/revenue-hack/go-practice/src/ch07/ex15/eval"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	var expr eval.Expr
	for {
		fmt.Println("----------------------")
		if !stdin.Scan() {
			fmt.Println("scan error")
		}
		expression := stdin.Text()
		fmt.Println(expression)
		var err error
		expr, err = eval.Parse(expression)
		if err != nil {
			fmt.Printf("parse error %v\n", expr)
			panic(err)
		} else {
			break
		}
	}
	// vars
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Println("Expr check error")
		panic(err)
	}
	// env
	var envs = eval.Env{}
	for v := range vars {
		fmt.Printf("express %v: ", v)
		if !stdin.Scan() {
			fmt.Println("scan error")
		}
		value, err := strconv.ParseFloat(stdin.Text(), 64)
		if err != nil {
			fmt.Println("parse float error")
		}
		envs[v] = value
	}
	fmt.Printf("answer %v = %v\n", expr, expr.Eval(envs))
}
