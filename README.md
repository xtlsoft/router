# Go-Router

Maybe the best router library in Golang.

Designed for all usages. Almost the same as XPHP's router.

## Installation

```bash
go get -u gopkg.in/xtlsoft/router.v1
```

## Quickstart

A Simple Example:

```go
package main

import (
    "fmt"
    router "gopkg.in/xtlsoft/router.v1"
)

type Response struct {
    StatusCode int
    Body string
}

func main() {

    r := router.New()

    r.Group("/", func (g *router.Group){
        g.On("GET", "test/{test}", func (req router.Request) router.Response {
            return &Response{
                StatusCode: 200,
                Body: "Hello World!",
            }
        })
    })

    fmt.Println( r.Handle("GET", "/test/abc", &router.DefaultRequest{}) )

}
```