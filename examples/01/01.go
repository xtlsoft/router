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
        g.On("GET", "test/{test}/{@int:another}", func (req router.Request) router.Response {
            return &Response{
                StatusCode: 200,
                Body: "Hello World!" + req.(*router.DefaultRequest).RouterVariable["test"],
            }
        })
    })

    fmt.Println( r.Handle("GET", "/test/abc/234", &router.DefaultRequest{}) )

}