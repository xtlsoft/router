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

func (this *Rule) Match(uri string) (matched bool, variables map[string]string, callback func (Request) Response) {

	matched, variables = CheckMatch(this.Parsed, uri)
	
	callback = this.Controller

	return

}