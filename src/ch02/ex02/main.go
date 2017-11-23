package main

import (
	"bufio"
	"ch02/ex02/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				panic(err)
			}
			temperature(t)
			weight(t)
			length(t)
		}
	} else {
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			if err := stdin.Err(); err != nil {
				fmt.Println("exit")
				os.Exit(1)
			}
			arg := stdin.Text()
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				panic(err)
			}
			temperature(t)
			weight(t)
			length(t)
		}
	}
}

func temperature(arg float64) {
	f := tempconv.Fahrenheit(arg)
	c := tempconv.Celsius(arg)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func weight(arg float64) {
	lb := tempconv.Lb(arg)
	kg := tempconv.Kg(arg)
	fmt.Printf("%s = %s, %s = %s\n", lb, tempconv.LbToKg(lb), kg, tempconv.KgToLb(kg))
}

func length(arg float64) {
	ft := tempconv.Ft(arg)
	m := tempconv.Meter(arg)
	fmt.Printf("%s = %s, %s = %s\n", ft, tempconv.FtToMeter(ft), m, tempconv.MeterToFt(m))
}
