package main

import (
	router "../.."
	"fmt"
)

func main(){

	parsed := router.ParseRule("/test/{module}/{[0-9]+:cid}/{[a-z]+:ident}.html")

	fmt.Println(parsed)

	fmt.Println(router.CheckMatch(parsed, "/test/aaa/10/afdf.html") )

}