package router

import (
	"strings"
)

type Router struct {
	
	groups []*Group
	handler func(bool, interface{}, interface{}, map[string]string) interface{}

}

func New() *Router {
	r := new(Router)
	r.SetHandler(DefaultHandler)
	return r
}

func (this *Router) Group(base string, callback func(*Group)) *Group {
	
	var grp = new(Group)

	grp.base = base
	grp.notFoundResponse = DefaultNotFoundController

	this.groups = append(this.groups, grp)

	callback(grp)

	return grp

}

func (this *Router) Parse(method string, uri string) (isMatched bool, controller interface{}, variables map[string]string) {

	method = strings.ToLower(method)

	var handleGroup *Group

	for _, v := range this.groups {
		if strings.HasPrefix(uri, v.base) {
			handleGroup = v
		}
	}

	return handleGroup.Parse(method, uri)

}

func (this *Router) Any(controller interface{}){
	DefaultNotFoundController = controller
}

func (this *Router) SetHandler(handler func(bool, interface{}, interface{}, map[string]string) interface{}) {

	this.handler = handler

}

func (this *Router) Handle(method string, uri string, request interface{}) (Response interface{}, isMatched bool) {
	
	h := this.handler
	is, con, vars := this.Parse(method, uri)

	resp := h(is, request, con, vars)
	
	return resp, is

}
