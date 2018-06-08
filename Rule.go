package router

import (

)

type Rule struct {

	Method string
	Uri string
	Controller func (Request) Response
	
	Parsed *ParsedRule

}

func (this *Rule) Parse() {
	
	this.Parsed = ParseRule(this.Uri)
	
}

func (this *Rule) Match() {

	

}