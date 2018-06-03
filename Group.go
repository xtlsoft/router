package router

import (

)

type Group struct {

	base string

	rules []*Rule

}

func (this *Group) On(method string, uri string, controller func(*Request) *Response) *Rule {

	rule := &Rule{
		Uri: uri,
		Controller: controller,
	}

	this.rules = append(this.rules, rule)

	return rule

}