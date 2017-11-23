package test

import (
	"ch02/ex02/tempconv"
	"strconv"
	"testing"
)

func TestTemperatureFahrenheit(t *testing.T) {
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

func TestTemperatureCelsius(t *testing.T) {
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

func TestWeightLb(t *testing.T) {
	lb := tempconv.Lb(10)
	if lb != 10 {
		t.Error("TestWeightLb func Lb not 10")
	}
	if tempconv.LbToKg(lb) != 4.54 {
		t.Error("TestWeightLb func LbToKg not 4.54")
	}
}

func TestWeightKg(t *testing.T) {
	kg := tempconv.Kg(10)
	if kg != 10 {
		t.Error("TestWeightKg func Kg not 10")
	}
	if tempconv.KgToLb(kg) != 22.05 {
		t.Error("TestWeightKg func KgToLb not 22.05")
	}
}

func TestLengthFt(t *testing.T) {
	ft := tempconv.Ft(10)
	if ft != 10 {
		t.Error("TestLengthFt func ft not 10")
	}
	if tempconv.FtToMeter(ft) != 32.8 {
		t.Error("TestLengthFt func FtToMeter not 32.8")
	}
}

func TestLengthMeter(t *testing.T) {
	m := tempconv.Meter(10)
	if m != 10 {
		t.Error("TestLengthMeter func m not 10")
	}
	if tempconv.MeterToFt(m) != 3.048780487804878 {
		t.Error("TestLengthMeter func MeterToFit not 3.048780487804878")
	}
}
