package test

import (
	"ch02/ex01/tempconv"
	"strconv"
	"testing"
)

func TestFahrenheit(t *testing.T) {
	temp, err := strconv.ParseFloat("10", 64)
	if err != nil {
		panic(err)
	}
	f := tempconv.Fahrenheit(temp)
	if f != 10 {
		t.Error("TestFahrenheit func f not 10")
	}
	if tempconv.FToC(f) != -12.222222222222221 {
		t.Error("TestFahrenheit func FToC not -12.222222222222221")
	}
}

func TestCelsius(t *testing.T) {
	temp, err := strconv.ParseFloat("10", 64)
	if err != nil {
		panic(err)
	}
	c := tempconv.Celsius(temp)
	if c != 10 {
		t.Error("TestFahrenheit func c not 10")
	}
	if tempconv.CToF(c) != 50 {
		t.Error("TestFahrenheit func CToF not 50")
	}
}
