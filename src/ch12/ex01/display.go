package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		var str bytes.Buffer
		str.WriteString(v.Type().String())
		str.WriteString("{")
		num := v.NumField()
		for i := 0; i < num; i++ {
			str.WriteString(v.Type().Field(i).Name + ":")
			str.WriteString(formatAtom(v.Field(i)))
			if i != num-1 {
				str.WriteString(",")
			}
		}
		str.WriteString("}")
		return str.String()
	case reflect.Array:
		var str bytes.Buffer
		str.WriteString(v.Type().String())
		str.WriteString("{")
		num := v.Len()
		for i := 0; i < num; i++ {
			str.WriteString(formatAtom(v.Index(i)))
			if i != num-1 {
				str.WriteString(",")
			}
		}
		str.WriteString("}")
		return str.String()
	default:
		return v.Type().String() + " value"
	}
}

func main() {
	type Test struct {
		name string
		num  int
	}
	testmap := map[Test]string{
		Test{name: "miyakawa", num: 2}: "rtc",
		Test{name: "golang", num: 2}:   "programing",
	}
	Display("testmap", testmap)

	testmap2 := map[[2]string]int{
		[2]string{"miyakawa", "rtc"}:      1,
		[2]string{"golang", "programing"}: 2,
	}
	Display("testmap2", testmap2)
}
