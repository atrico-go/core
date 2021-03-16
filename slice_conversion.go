package core

import (
	"errors"
	"fmt"
	"reflect"
)

// Convert slice of one type to another
// Types must be convertible
func ConvertSlice(input interface{}, output interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	iVal := reflect.ValueOf(input)
	if iVal.Kind() != reflect.Slice {
		return errors.New("input type not a slice")
	}
	oVal := reflect.ValueOf(output)
	if oVal.Kind() != reflect.Ptr || oVal.Elem().Kind() != reflect.Slice {
		return errors.New("output type not a ptr to a slice")
	}
	iLen := iVal.Len()
	outputV := reflect.MakeSlice(oVal.Elem().Type(), iLen, iLen)
	for i := 0; i < iLen; i++ {
		outputV.Index(i).Set(reflect.ValueOf(iVal.Index(i).Interface()))
	}
	oVal.Elem().Set(outputV)
	return nil
}
