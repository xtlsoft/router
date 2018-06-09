package main

import (
    "fmt"
    router "../.."
)

func main() {

    r := router.New()

    router.AddShortcut("@test", "i[0-9]+");

    r.Group("/", func (g *router.Group){
        g.On("GET", "test/{test}/{@test:another}", func (req router.Request) router.Response {
            return &router.DefaultResponse{
                StatusCode: 200,
                Body: "Hello World!" + req.(*router.DefaultRequest).RouterVariable["test"],
            }
        })
        g.Any(func (router.Request) router.Response{
            return &router.DefaultResponse{
                Body: "404 Not Found",
            }
        })
    })

    fmt.Println( r.Handle("GET", "/test/abc/i234", &router.DefaultRequest{}) )

}