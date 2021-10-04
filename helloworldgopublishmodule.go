package helloworldgopublishmodule

import (
	"fmt"
)

const (
	ModuleName = "github.com/parkgang/helloworld-go-publish-module"
)

// 외부로 잘 노출되는 것을 확인할 수 있습니다.
func PublicPrintModuleName() {
	fmt.Printf("Public Func: %s\n", ModuleName)
}

// 외부로 노출되지 않는 것을 확인할 수 있습니다.
func privatePrintModuleName() {
	fmt.Printf("Private Func: %s\n", ModuleName)
}

// 외부로 노출되지 않지만 공개되는 함수에 의해 호출되는 것을 확인합니다.
func ExpandFunc() {
	privatePrintModuleName()
}
