# unsafe

[go1.21rc2/src/unsafe](https://github.com/golang/go/tree/go1.21rc2/src/unsafe)

unsafeパッケージは `unsafe.go` ファイルだけで構成されているパッケージで、

> Packages that import unsafe may be non-portable and are not protected by the Go 1 compatibility guidelines.

> Packages that import unsafe may depend on internal properties of the Go implementation. We reserve the right to make changes to the implementation that may break such programs.

コメントアウトや[Go 1 and the Future of Go Programs](https://go.dev/doc/go1compat)の上記にも記載されているように、 unsafeパッケージを使用したコードの互換性は担保されないことが明示されているので、基本的にプロダクトだったりで直接利用することは避けるべきパッケージらしい。

また、unsafe.go自体には実装は記述されていない。

## 使い方

`unsafe.Pointer` を使用することで、下記のように任意の型のポインター型を別の型のポインター型に変更することができる。

```go
package main

import (
	"fmt"
	"unsafe"
)

type StructA struct {
	Name string
	Age  int
}

type StructB struct {
	Msg string
	ID  int
}

func main() {
	var sa = StructA{"sa", 20}
	var sb = *(*StructB)(unsafe.Pointer(&sa))

	fmt.Printf("%#v\n", sa)
	// main.StructA{Name:"sa", Age:20}

	fmt.Printf("%#v\n", sb)
	// main.StructB{Msg:"sa", ID:20}
}
```

https://go.dev/ref/spec#Package_unsafe
