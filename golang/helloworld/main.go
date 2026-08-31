// package는 main으로 시작하여 컴파일이 가능한 코드를 작성할 수 있다.
// 그 외 패키지는 라이브러리 함수 용도나 객체로서 사용된다.
package main

// Import는 내부 라이브러이와 외부 라이브러리를 가져와서 사용할 수 있다.
// 하지만 외부 라이브러리는 go.mod와 go.sum에 등록되어야 하며, github 리포지토리에서 관리한다.
import (
	"fmt"
	// "os"
	"time"
)

// main 함수는 프로그램의 시작점으로 main.go 파일에 반드시 하나의 함수로만 이루어져 있으며,
// package main으로 시작되는 파일에서만 정의할 수 있다.
func main() {
	fmt.Println("Hello World " + time.Now().String())
}
