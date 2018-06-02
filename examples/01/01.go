package main

import (
    "fmt"
    router "../.."
)

func main() {

    r := router.New()

    r.Group("/", func (r){
        return r
    })

}