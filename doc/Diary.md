# RachelHomeService with echo

## 첫째날

전체적인 얼개는

1. 서비스 객체를 생성한 뒤
2. 각 기능들을 등록하고
3. 서비스를 실행.

세 단계에 걸쳐서 구성해보려고 시도해 보았다.

### 모듈 만들기

제일 먼저 모듈 선언

```bash
go mod init github.com/gwangyi/RachelHomeService
```

go 1.12부터 추가된 `go mod` 기능은 좀 더 말이 되는 형태로 go 모듈들 간의 의존성을 관리할 수 있게 해준다.

go  객체의 특징 중 하나는, 다른 언어처럼 클래스를 구성하는 요소들을 반드시 한 곳에 모아 놓을 필요가 없다는 것이다.
필드들은 한 곳에 모아야 하지만, 메소드에 해당하는 Receiver Function 들은 같은 패키지 내에만 있으면 큰 문제가 없다.
인터페이스를 정의할 때는 그렇지 않다. 큰 문제가 없다는 것은 함수의 정의에만 해당한다.

각 기능들별로 별도의 소스파일로 분리한 뒤 등록하는 함수를 서비스의 메소드로 선언할 계획이다.

### 의존성 설치

pip나 yarn 등과는 달리, go mod를 사용할 경우 대충 코드를 작성한 뒤에 의존성을 설치해주는 명령을 실행하면 의존성을 모아서 알아서 설치해준다.
아마도 import 구문 안에 이미 의존성에 대한 정보가 다 적혀 있기 때문인 것 같다.

```bash
go mod vendor
```

새로운 의존성이 필요할 때도 일단 import 부터 한 뒤에 `go mod vendor` 를 실행해 주면 필요한걸 내려받는다.

### echo

```go
e := echo.New()
e.GET("/endpoint", func(c echo.Context) error {
	c.JSON(http.StatusOK, object)
	return nil
})
```

대략 이렇게 handler들을 등록하면 된다.
끌 때는 `e.Shutdown()` 함수를 쓸 수 있는데, 패러미터로 context를 받는다. 이런 방식이 go에서는 종종 있는 듯 하다.
뭘 써야될지 모를 때는 `context.TODO()` 를 쓰면 되고, Top-level context 가 필요할 때는 `context.Background()` 를 쓰면 된다고 한다. 
context를 잘 쓰면 deadline, timeout, cancel 등을 표준적인 형태로 다룰 수 있는 모양이다.


### 오늘의 참고자료

* [go mod][1]
* [echo][2], [echo godoc][3]
* [context][4]

[1]: https://blog.golang.org/using-go-modules
[2]: https://echo.labstack.com/
[3]: https://pkg.go.dev/github.com/labstack/echo?tab=doc
[4]: https://pkg.go.dev/context?tab=doc
