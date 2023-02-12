# GraphQL sample started

## Exsample command

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

2023/01/31 09:38:24 connect to `http://localhost:8080/` for GraphQL playground

### Description

- schema

GraphQL の設定ファイルです。
配列で複数指定することもできますが、 Glob で指定できるため schema/*.graphql のようにディレクトリ配下の設定を一括取り込みしたりもできます。

- exec

後述する models_gen.go をどのようなファイル名、どのようなパッケージ名で出力するか設定できます。

- filename にはパスから指定できるのでディレクトリも変更できます。

- package を指定すれば出力するファイルのパッケージ名を変更できます。

出力するディレクトリ名と合わせる必要があります

- model

後述する generated.go をどのようなファイル名、どのようなパッケージ名で出力するか設定できます。 (設定方法は execと同じため省略)

- models

手動で生成したモデルを gqlgen に教えてあげるための設定です。 {親ディレクトリまでのパス}/{パッケージ名}.{モデル名} という具合に指定します。 詳しくは次のセクションまで読み進めてください。

- resolver

後述する resolver.go をどのようなファイル名、どのようなパッケージ名で出力するか設定できます。 (設定方法は execと同じため省略)
type はリゾルバに設定する名前です。変えたからと言って挙動が変わったりはしません。 通常は Resolver のままにしておくのがよいでしょう。

- autobind

手動で作成したモデルを自動的に読み込むための設定です。 配列でパスを指定すれば該当するモデルを読み込んでくれます。
v0.9.1 から使えます。

How to configure gqlgen using gqlgen.yml - gqlgen
