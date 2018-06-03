package main

import (
    "fmt"
    router "../.."
)

type Response struct {
    StatusCode int
    Body string
}

func main() {

    r := router.New()

    r.Group("/", func (g *router.Group){
        g.On("GET", "test/{test}", func (req *router.Request) resp *Response {
            return &Response{
                StatusCode: 200,
                Body: "Hello World!",
            }
        })
    })

    r.Parse()

    println( r.Handle("GET", "/test/abc", &router.Request{}) )

}