package router

import (
	"strings"
)

type Group struct {

	base string

	rules []*Rule

	notFoundResponse interface{}

}

func (this *Group) On(method string, uri string, controller interface{}) *Rule {

	rule := &Rule{
		Method: strings.ToLower(method),
		Uri: uri,
		Controller: controller,
	}

	rule.Parse()

	this.rules = append(this.rules, rule)
	
	return rule

}

func (this *Group) Parse(method string, uri string) (isMatched bool, controller interface{}, variables map[string]string) {

	var flag = false
	var vars = map[string]string{}
	var ctrler interface{}

	for _, v := range this.rules {
		if v.Method == "*" || v.Method == method {
			matched, variable, callback := v.Match(uri)
			if matched {
				flag = true
				vars = variable
				ctrler = callback
			}
		}
	}

	if !flag {
		return false, this.notFoundResponse, map[string]string{"Uri": uri, "Method": method}
	}

	return true, ctrler, vars

}

func (this *Group) Any(controller interface{}) {

	this.notFoundResponse = controller

}