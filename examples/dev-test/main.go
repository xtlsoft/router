package main

import (
	router "../.."
	"fmt"
)

func main(){

	fmt.Print( router.ParseRule("/.*2738@#34/{module}/{[0-9]:cid}/{[a-z]:ident}.html") )

}