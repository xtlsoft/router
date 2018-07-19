package router

import (
	"reflect"
)

func ExportReflect(t interface{}) {

	r := reflect.ValueOf(t)
	// No time to update, so just so so...
	_ = r
	
}
