
## Dive into Golang 2.0
## Task 1

```
go mod init <NAME>
```

```
go test -v -coverpkg local/pkg -coverprofile cover.out .
go tool cover -html cover.out -o cover.html
open cover.html
```

```
go run .
--- Dive into Golang 2.0 ---
"zzzzcccUUUuu" > "U3c3u2z4"
"aaabbbccccc" > "a3b3c5"
"cccccaaabbb" > "a3b3c5"
```

### Примечание:
Псевдоним `str` для базового типа `string` по аналогии `int` - это сокращенное слово `integer`
```
// alias declaration
type str = string

var i int
var s str
```