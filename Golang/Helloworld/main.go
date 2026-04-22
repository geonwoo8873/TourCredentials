//// N.1
// `package main`은 컴파일 시작점을 포함하는 기본 패키지이다. 패키지는 패키지의 이름으로 `main`으로 사용할 수 없다.
// `main()` 함수가 없기 때문에 실행 파일을 생성할 수 없고, 다른 패키지에서 외부 패키지로 사용된다.
//// N.2
// `import`는 다른 패키지를 현재 패키지로 가져오는 데 사용된다.
// `fmt` 패키지는 포맷된 입출력 함수를 구현하는 Go 언어 표준 라이브러리 패키지이다.
//// N.3
// func <function name>(value type) <return type> {...}

// N.1
package main

// N.2
import "fmt"

// N.3
func Calculate(t int, p float64) float64 {
	return float64(t) * p * 10000
}

func main() {
	fmt.Println(Calculate(45, 0.5))
}
