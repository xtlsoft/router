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

func main() {

    r := router.New()

    r.Group("/", func (r){
        return r
    })

}
```