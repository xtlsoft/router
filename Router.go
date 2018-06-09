package router

import (
	"strings"
)

type Router struct {
	
	groups []*Group

}

func New() *Router {
	return new(Router)
}

func (this *Router) Group(base string, callback func(*Group)) *Group {
	
	var grp = new(Group)

	grp.base = base
	grp.notFoundResponse = DefaultNotFoundController

	this.groups = append(this.groups, grp)

	callback(grp)

	return grp

}

func (this *Router) Handle(method string, uri string, request Request) (resp Response, isMatched bool) {

	var handleGroup *Group

	for _, v := range this.groups {
		if strings.HasPrefix(uri, v.base) {
			handleGroup = v
		}
	}

	return handleGroup.Handle(method, uri, request)

}

func (this *Router) Any(controller func(Request) Response){
	DefaultNotFoundController = controller
}