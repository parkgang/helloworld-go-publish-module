package main

import (
	"fmt"

	"github.com/parkgang/helloworld-go-publish-module/libs"
)

const (
	ModuleName = "github.com/parkgang/helloworld-go-publish-module"
)

func PublicPrintModuleName() string {
	return ModuleName
}

func privatePrintModuleName() string {
	return ModuleName
}

func main() {
	fmt.Printf("hello, %s\n", privatePrintModuleName())
	fmt.Println(libs.OtherPrint())
}
