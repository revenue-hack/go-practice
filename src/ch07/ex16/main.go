package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/revenue-hack/go-practice/src/ch07/ex16/eval"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	exprParam, varsParam := getParameter(r)
	if exprParam != "" {
		expr, err := eval.Parse(exprParam)
		if err != nil {
			fmt.Printf("parse error %v\n", expr)
			panic(err)
			return
		}
		vars := make(map[eval.Var]bool)
		if err := expr.Check(vars); err != nil {
			fmt.Println("Expr check error")
			panic(err)
			return
		}
		var envs = eval.Env{}
		for v := range vars {
			if !inKey(varsParam, string(v)) {
				fmt.Fprintf(w, "expression key nothing %v\n", v)
				return
			}
			envs[v] = varsParam[string(v)]
		}
		fmt.Fprintf(w, "answer %v = %v\n", expr, expr.Eval(envs))
	} else {
		fmt.Fprintln(w, "answer nothing because expr parameter nothing")
	}
}

func inKey(array map[string]float64, key string) bool {
	for v := range array {
		if v == key {
			return true
		}
	}
	return false
}

func getParameter(r *http.Request) (string, map[string]float64) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	var expr string
	var vars = make(map[string]float64)
	for k, v := range r.Form {
		if k == "expr" {
			expr = v[0]
		} else {
			value, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err)
			}
			vars[k] = value
		}
	}
	return expr, vars
}
