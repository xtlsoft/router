package router

import (
	
)

var DefaultNotFoundResponse = &HttpResponse{
		StatusCode: 404,
		Body: `<!Doctype html>
<html>
 <head>
  <meta charset="utf-8" />
 </head>
 <body>
  <h1 align="center">HTTP 404</h1>
  <p align="center">Not Found</p>
  <hr />
  <p><i>Powered by xtlsoft/router (NotFound Module)</i></p>
 </body>
</html>`,
}

var DefaultNotFoundController interface{} = func(req Request) Response {
	return DefaultNotFoundResponse
}

var Shortcuts = map[string]string{
	"@int" : "[0-9]+",
	"@ident" : "[A-Za-z0-9_-]+",
}

func AddShortcut(alias string, regex string){
	Shortcuts[alias] = regex
}

var DefaultHandler  = func(isMatched bool, request interface{}, controller interface{}, variables map[string]string) (response interface{}) {

	r := request.(Request)
	
	c := controller.(func (Request) Response)

	r.SetRouterVariable(variables)

	return c(r)

}

func SetHandler(handler func(bool, interface{}, interface{}, map[string]string) interface{}) {

	DefaultHandler = handler

}