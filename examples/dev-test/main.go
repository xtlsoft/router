package main

import (
	router "../.."
	"fmt"
)

func main(){

	fmt.Print( router.ParseRule("/test/{module}/{[0-9]:cid}/{[a-z]:ident}.html") )

}