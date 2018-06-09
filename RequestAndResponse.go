package router

import (

)

type Request interface {
	SetRouterVariable(map[string]string)
}

type Response interface {
	
}

type DefaultRequest struct {
	RouterVariable map[string]string
}

type DefaultResponse struct {
	
}

func (this *DefaultRequest) SetRouterVariable(vars map[string]string){
	this.RouterVariable = vars
}