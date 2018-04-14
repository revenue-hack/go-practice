package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func Pack(ptr interface{}, domain string) (string, error) {
	var buf bytes.Buffer
	buf.WriteString(domain + "?")
	values := reflect.ValueOf(ptr).Elem()
	for i := 0; i < values.NumField(); i++ {
		if i != 0 {
			buf.WriteString("&")
		}
		key := getKey(i, values)
		err := appendParams(&buf, values.Field(i), key)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func getKey(i int, values reflect.Value) string {
	key := values.Type().Field(i).Tag.Get("http")
	if key == "" {
		key = strings.ToLower(values.Type().Field(i).Name)
	}
	return key
}

func appendParams(buf *bytes.Buffer, value reflect.Value, key string) error {
	switch value.Kind() {
	case reflect.Int:
		fmt.Fprintf(buf, "%s=%d", key, value.Int())
		return nil
	case reflect.String:
		fmt.Fprintf(buf, "%s=%s", key, value.String())
		return nil
	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if err := appendParams(buf, value.Index(i), key); err != nil {
				return err
			}
			if i != value.Len()-1 {
				buf.WriteString("&")
			}
		}
		return nil
	case reflect.Bool:
		fmt.Fprintf(buf, "%s=%v", key, value.Bool())
		return nil
	default:
		panic(fmt.Sprintf("unsupported kind %v", value.Kind()))
	}
	return nil
}
