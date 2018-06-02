package router

import (

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

	this.groups = append(this.groups, grp)

	callback(grp)

	return grp

}