package main

import (
    "fmt"
    router "../.."
)

func main() {

    r := router.New()

    r.Group("/", func (g *router.Group){
        g.On("GET", "/", ".")
    })

}