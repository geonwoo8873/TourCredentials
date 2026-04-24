# SetUp Project File

## Install Go

```ps
winget install --id=GoLang.Go -e
```

## Tutorial create project

### Create module
```bash
cd $PROJECT$
go mod init $PROJECT$/example
```

### Create and Edit codefile
```go
// Path : $PROJECT$/main.go

// N.1
package main

// N.2
import "fmt"

func Calculate(t int, p float64) float64 {
	return float64(t) * p * 10000
}

// N.3
func main() {
	fmt.Println(Calculate(45, 0.5))
}
```

* N.1
  * `package main`은 컴파일 시작점을 포함하는 기본 패키지이다. 패키지는 패키지의 이름으로 `main`으로 사용할 수 없다.
  * `main()` 함수가 없기 때문에 실행 파일을 생성할 수 없고, 다른 패키지에서 외부 패키지로 사용된다.
* N.2
  * `import`는 다른 패키지를 현재 패키지로 가져오는 데 사용된다.
  * `fmt` 패키지는 포맷된 입출력 함수를 구현하는 Go 언어 표준 라이브러리 패키지이다.
* N.3
  * func <function name>(value type) <return type> {...}
* N.4
  * `var` or `const`는 변수 선언 키워드로 var의 경우 값을 업데이트[mutable]이 가능하며,
  * `const`는 상수 선언 키워드로 const의 경우 값을 업데이트[immutable]할 수 없다.

### Go build architecture

```bash
# CMD
go tool dist list
```

```bash
# Result
aix/ppc64
{...}
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/loong64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
{...}
windows/386
windows/amd64
windows/arm64
```

> [!TIP]
> 사용자의 사용하는 칩셋이나 개발 환경에서 빌드하고자 하는 실행 파일을 만들 때 활용하면 유용하다

#### Go select of build architecture

```bash
# CMD
GOOS=<OS> GOARCH=<ChipSet> go buuild
```

```bash
# Example
GOOS=linux GOARCH=arm64 go build
```