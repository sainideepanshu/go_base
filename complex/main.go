package main

//relative import path are not supported in Go modules, so this will not work in a module context
//Capital letters work only cross-package

// In same whether capital or small letter, it will not matter
// the issue is that we are not compiling the lib.go file
// to resolve the issue we should run go run *.go
// in windows go run *.go will not work, so we need to specify the file names
// in linux it will work with go run *.go	
// go run .

// import (
	
// 	"fmt"
// 	app "../app"
// )

import (
	
	"fmt"
	app "GoBase/app"
)

func main() {
	s := hello1()
	fmt.Println(s)
	c := app.NewComplex(1.0, 2.0)
	fmt.Println(c.Real, c.Imaginary)
}