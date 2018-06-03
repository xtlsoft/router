package main

import (
	router "../.."
	"fmt"
)

func main(){

	fmt.Print( router.ParseRule("/{aaa}/{nnn}") )

}