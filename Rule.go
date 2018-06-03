package router

import (

)

type Rule struct {

	Uri string
	Controller func (Request) Response
	
	Parsed *ParsedRule

}

func (this *Rule) Parse() {
	
	this.Parsed = ParseRule(this.Uri)
	
}
