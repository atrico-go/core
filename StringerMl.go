package core

import (
	"fmt"
	"io"
	"os"
)

// Multiline version of Stringer
type StringerMl interface {
	// Representation of object as multiple lines
	StringMl(params ...interface{}) []string
}

func DisplayMultiline(obj StringerMl, params ...interface{}) {
	FdisplayMultiline(os.Stdout, obj, params...)
}

func FdisplayMultiline(writer io.Writer, obj StringerMl, params ...interface{}) {
	for _, line := range obj.StringMl(params...) {
		fmt.Fprintln(writer, line)
	}
}
