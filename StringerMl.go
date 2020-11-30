package core

import "fmt"

// Multiline version of Stringer
type StringerMl interface {
	// Representation of object as multiple lines
	StringMl(params ...interface{}) []string
}

func DisplayMultiline(obj StringerMl, params ...interface{}) {
	for _,line := range obj.StringMl(params...) {
		fmt.Println(line)
	}
}