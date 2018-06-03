package router

import (

)

type Rule struct {

	Uri string
	Controller func (*Request) *Response

}
