package main

import (
    "fmt"
    router "../.."
)

func main() {

    r := router.New()

    router.AddShortcut("@test", "i[0-9]+");

    fmt.Println("Server started.")

    r.Group("/", func (g *router.Group){
        g.On("GET", "test/{test}/{@test:another}", func (req router.Request) router.Response {
            return &router.HttpResponse{
                StatusCode: 200,
                Body: "Hello World!" + fmt.Sprintln(req),
            }
        })
    })

	// fmt.Println( r.Handle("GET", "/test/abc/i234", &router.DefaultRequest{}) )
	
	router.HttpServe(r, "127.0.0.1:28080")

}