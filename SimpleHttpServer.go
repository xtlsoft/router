package router

import (	
	"net/http"
)

type HttpRequest struct {
	*http.Request
	RouterVariable map[string]string
}

type IHttpResponse interface {
	GetStatusCode() int
	GetBody() string
}

type HttpResponse struct {
	StatusCode int
	Body string
}

func (this *HttpResponse) GetStatusCode() int {
	return this.StatusCode
}

func (this *HttpResponse) GetBody() string {
	return this.Body
}

func HttpServe(router *Router, address string) {

	http.HandleFunc("/", func (rw http.ResponseWriter, req *http.Request) {
		me := req.Method
		ur := req.URL
		var request *HttpRequest
		request = &HttpRequest{}
		request.Request = req
		resp, _ := router.Handle(me, ur.RequestURI(), request)

		rw.WriteHeader(resp.(IHttpResponse).GetStatusCode())
		rw.Write([]byte(resp.(IHttpResponse).GetBody()))
	})

	http.ListenAndServe(address, nil)

}

func (this *HttpRequest) SetRouterVariable(vars map[string]string){
	this.RouterVariable = vars
}