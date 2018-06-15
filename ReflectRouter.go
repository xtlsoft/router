package router

import (
	"reflect"
)

func ExportReflect(t interface{}) {

	r := reflect.ValueOf(t)

}