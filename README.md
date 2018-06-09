# Go-Router

Maybe the best router library in Golang.

Designed for all usages. Almost the same as XPHP's router.

## Installation

```bash
go get -u gopkg.in/xtlsoft/router.v1
```

## Quickstart

A Simple Example: (In `/examples/01/01.go`)

```go
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
```

## Documention

The router is very easy to use.
The document is avaliable in `Document.md`.