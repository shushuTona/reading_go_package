package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/token"
)

func main() {
	fset := token.NewFileSet()

	src1 := "var a = 100"
	src2 := "func b () int { return 200 }"
	src3 := "var c = []int{0, 1, 2}"

	fset.AddFile("file_1", -1, len(src1))
	fset.AddFile("file_2", -1, len(src2))
	fset.AddFile("file_3", -1, len(src3))

	var buf bytes.Buffer
	encode := func(x any) error {
		return json.NewEncoder(&buf).Encode(x)
	}

	fset.Write(encode)

	fmt.Println(buf.String())

	fmt.Printf("before fset.Base() : %#v\n", fset.Base())
	fmt.Printf("before fset.File(1) : %#v\n", fset.File(1))
	fmt.Printf("before fset.File(13) : %#v\n", fset.File(13))
	fmt.Printf("before fset.File(42) : %#v\n", fset.File(50))
	// before fset.Base() : 65
	// before fset.File(1) : &token.File{name:"file_1", base:1, size:11, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
	// before fset.File(13) : &token.File{name:"file_2", base:13, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
	// before fset.File(42) : &token.File{name:"file_3", base:42, size:22, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}

	jstr := `{
    "Base": 42,
    "Files": [
        {
            "Name": "file_1",
            "Base": 1,
            "Size": 11,
            "Lines": [
                0
            ],
            "Infos": null
        },
        {
            "Name": "file_2",
            "Base": 13,
            "Size": 28,
            "Lines": [
                0
            ],
            "Infos": null
        }
    ]
}`
	bufR := bytes.NewBufferString(jstr)
	decode := func(x any) error {
		return json.NewDecoder(bufR).Decode(x)
	}
	fset.Read(decode)

	fmt.Printf("after fset.Base() : %#v\n", fset.Base())
	fmt.Printf("after fset.File(1) : %#v\n", fset.File(1))
	fmt.Printf("after fset.File(13) : %#v\n", fset.File(13))
	fmt.Printf("after fset.File(42) : %#v\n", fset.File(50))
	// after fset.Base() : 42
	// after fset.File(1) : &token.File{name:"file_1", base:1, size:11, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
	// after fset.File(13) : &token.File{name:"file_2", base:13, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
	// after fset.File(42) : (*token.File)(nil)
}
