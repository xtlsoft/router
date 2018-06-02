package router

import (

)

type Rule struct {

	Rule string
	Controller func (Request) Response

}
