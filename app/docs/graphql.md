# GraphQL sample started

### 1

`go get github.com/99designs/gqlgen`

`mkdir tools && cd tools`

`touch tools.go`

- /tools/tools.go

```go
package tools
import _ “github.com/99designs/gqlgen”
```

### 2

`$ go run github.com/99designs/gqlgen init`

Creating gqlgen.yml

Creating graph/schema.graphqls

Creating server.go

Generating...

go: downloading github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883

go: downloading github.com/arbovm/levenshtein v0.0.0-20160628152529-48b4e1c0c4d0

go: downloading github.com/dgryski/trifles v0.0.0-20200323201526-dd97f9abfb48

go: downloading github.com/gorilla/websocket v1.5.0

go: downloading github.com/hashicorp/golang-lru v0.5.4

go: downloading github.com/sergi/go-diff v1.1.0

Exec "go run ./server.go" to start GraphQL server

### 3

`$ go run github.com/99designs/gqlgen generate`

### 4

`$ go run ./server.go`

2023/01/31 09:38:24 connect to http://localhost:8080/ for GraphQL playground
