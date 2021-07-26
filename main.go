package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type City struct {
	Name       string `json:"name"`
	Population int64  `json:"population"`
	GDP        int64  `json:"gdp,omitempty"`
	Mayor      string `json:"mayor"`
}

func main() {
	var u = User{"bob", 10}

	res, err := JSONEncode(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	c := City{"sf", 5000000, 0, ""}
	res, err = JSONEncode(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func JSONEncode(v interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	values := reflect.ValueOf(v)
	fields := reflect.TypeOf(v)
	fmt.Println()
	if values.Kind() != reflect.Struct {
		return nil, errors.New("the argument is not structure")
	}
	buf.WriteString("{\n")
	num := fields.NumField()
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		tag := field.Tag.Get("json")
		switch value.Interface().(type) {
		case string:
			if !(value.String() == "" && tag != "" && strings.Contains(tag, "omitempty")) {
				buf.WriteString("\t" + field.Name + ": \"" + value.String() + "\"\n")
			}
		case int64:
			if !(value.Int() == 0 && tag != "" && strings.Contains(tag, "omitempty")) {
				buf.WriteString("\t" + field.Name + ": " + strconv.FormatInt(value.Int(), 10) + "\n")
			}
			fmt.Println(tag)
		default:
			return nil, errors.New("wrong field type")
		}
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}
