package router

import (

)

type Group struct {

	base string

	rules []*Rule

	notFoundResponse func(Request) Response

}

func (this *Group) On(method string, uri string, controller func(Request) Response) *Rule {

	rule := &Rule{
		Method: method,
		Uri: uri,
		Controller: controller,
	}

	rule.Parse()

	this.rules = append(this.rules, rule)
	
	return rule

}

func (this *Group) Handle(method string, uri string, req Request) (resp Response, isMatched bool) {

	var flag = false
	var vars = map[string]string{}
	var controller func (Request) Response

	for _, v := range this.rules {
		if v.Method == "*" || v.Method == method {
			matched, variables, callback := v.Match(uri)
			if matched {
				flag = true
				vars = variables
				controller = callback
			}
		}
	}

	if !flag {
		return this.notFoundResponse(req), false
	}

	var response Response

	req.SetRouterVariable(vars)

	response = controller(req)

	return response, true

}

func (this *Group) Any(controller func(Request) Response) {

	this.notFoundResponse = controller

}