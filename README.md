# Overview

> my first go module publish

내가 만든 go module이 어떤 식으로 `publish` 되는지 확인하기 위한 repo 입니다.

# 테스트 방법

local에 go project를 생성하고 `go get github.com/parkgang/helloworld-go-publish-module` 으로 해당 module를 다운받아서 사용해보세요.

```go
package main

import (
	"fmt"

	helloworldgopublishmodule "github.com/parkgang/helloworld-go-publish-module"
	"github.com/parkgang/helloworld-go-publish-module/libs"
)

func main() {
	fmt.Printf("내가 만든 모듈에서 호출하는 함수와 변수\n")
	fmt.Printf("ModuleName: %s\n", helloworldgopublishmodule.ModuleName)
	helloworldgopublishmodule.PublicPrintModuleName()
	helloworldgopublishmodule.ExpandFunc()
	fmt.Printf("%s\n", libs.OtherPrint())
}
```
