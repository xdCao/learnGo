package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(joinStrings("hi", " i'm", " xdCao"))
	fmt.Println(printTypeValue(100, true, "hahha"))
}

func joinStrings(slist ...string) string {
	var b bytes.Buffer
	for _, s := range slist {
		b.WriteString(s)
	}
	return b.String()
}

func printTypeValue(slist ...interface{}) string {
	var b bytes.Buffer
	for _, s := range slist {
		str := fmt.Sprintf("%v", s)
		var typeString string
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}
		b.WriteString("value : ")
		b.WriteString(str)
		b.WriteString(" type: ")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}
