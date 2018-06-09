package router

import (

)

/*
Generated with "Generate_GroupAlias.php"
*/

func (this *Group) Get(uri string, controller func(Request) Response) *Rule {
	return this.On("Get", uri, controller)
}

func (this *Group) Post(uri string, controller func(Request) Response) *Rule {
	return this.On("Post", uri, controller)
}

func (this *Group) Put(uri string, controller func(Request) Response) *Rule {
	return this.On("Put", uri, controller)
}

func (this *Group) Delete(uri string, controller func(Request) Response) *Rule {
	return this.On("Delete", uri, controller)
}

func (this *Group) Patch(uri string, controller func(Request) Response) *Rule {
	return this.On("Patch", uri, controller)
}

func (this *Group) CLI(uri string, controller func(Request) Response) *Rule {
	return this.On("CLI", uri, controller)
}

func (this *Group) WebSocket(uri string, controller func(Request) Response) *Rule {
	return this.On("WebSocket", uri, controller)
}

