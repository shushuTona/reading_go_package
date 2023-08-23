package main

import (
	"fmt"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	fmt.Printf("fset : %#v\n\n", fset)
	// fset : &token.FileSet{mutex:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:atomic.Int32{_:atomic.noCopy{}, v:0}, readerWait:atomic.Int32{_:atomic.noCopy{}, v:0}}, base:1, files:[]*token.File(nil), last:atomic.Pointer[go/token.File]{_:[0]*token.File{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(nil)}}

	fmt.Printf("file.Base() : %#v\n\n", fset.Base())
	// file.Base() : 1

	srca := `package main

import (
	"fmt"
)

func main() {
	fmt.Println(100)
}
`
	fmt.Printf("len(srca) : %#v\n", len(srca))
	// len(srca) : 67

	file := fset.AddFile("a", -1, len(srca))

	fmt.Printf("file : %#v", file)
	// file : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}

	fmt.Printf("fset : %#v\n\n", fset)
	// fset : &token.FileSet{mutex:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:atomic.Int32{_:atomic.noCopy{}, v:0}, readerWait:atomic.Int32{_:atomic.noCopy{}, v:0}}, base:69, files:[]*token.File{(*token.File)(0xc000114120)}, last:atomic.Pointer[go/token.File]{_:[0]*token.File{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(0xc000114120)}}

	fmt.Printf("file.Base() : %#v\n\n", fset.Base())
	// file.Base() : 69

	srcb := "func b () int { return 200 }"
	fmt.Printf("len(srcb) : %#v\n", len(srcb))
	// len(srcb) : 28

	fset.AddFile("b", -1, len(srcb))
	fmt.Printf("fset : %#v\n\n", fset)
	// fset : &token.FileSet{mutex:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:atomic.Int32{_:atomic.noCopy{}, v:0}, readerWait:atomic.Int32{_:atomic.noCopy{}, v:0}}, base:98, files:[]*token.File{(*token.File)(0xc000114120), (*token.File)(0xc000114180)}, last:atomic.Pointer[go/token.File]{_:[0]*token.File{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(0xc000114180)}}

	fmt.Printf("file.Base() : %#v\n\n", fset.Base())
	// file.Base() : 98

	srcc := "c := 300"
	fmt.Printf("len(srcc) : %#v\n", len(srcc))
	// len(srcc) : 8

	fset.AddFile("c", -1, len(srcc))
	fmt.Printf("fset : %#v\n\n", fset)
	// fset : &token.FileSet{mutex:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:atomic.Int32{_:atomic.noCopy{}, v:0}, readerWait:atomic.Int32{_:atomic.noCopy{}, v:0}}, base:107, files:[]*token.File{(*token.File)(0xc000114120), (*token.File)(0xc000114180), (*token.File)(0xc0001141e0)}, last:atomic.Pointer[go/token.File]{_:[0]*token.File{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(0xc0001141e0)}}

	fmt.Printf("file.Base() : %#v\n\n", fset.Base())
	// file.Base() : 107

	for i := 0; i <= fset.Base(); i++ {
		pos := token.Pos(i)
		fmt.Printf("fset.File(%03d) : %#v\n", pos, fset.File(pos))
		// fset.File(000) : (*token.File)(nil)
		//
		// fset.File(001) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(002) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(003) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(004) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(005) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(006) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(007) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(008) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(009) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(010) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(011) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(012) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(013) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(014) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(015) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(016) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(017) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(018) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(019) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(020) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(021) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(022) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(023) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(024) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(025) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(026) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(027) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(028) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(029) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(030) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(031) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(032) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(033) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(034) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(035) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(036) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(037) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(038) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(039) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(040) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(041) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(042) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(043) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(044) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(045) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(046) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(047) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(048) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(049) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(050) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(051) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(052) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(053) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(054) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(055) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(056) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(057) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(058) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(059) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(060) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(061) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(062) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(063) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(064) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(065) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(066) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(067) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(068) : &token.File{name:"a", base:1, size:67, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		//
		// fset.File(069) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(070) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(071) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(072) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(073) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(074) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(075) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(076) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(077) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(078) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(079) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(080) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(081) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(082) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(083) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(084) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(085) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(086) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(087) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(088) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(089) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(090) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(091) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(092) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(093) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(094) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(095) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(096) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(097) : &token.File{name:"b", base:69, size:28, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		//
		// fset.File(098) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(099) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(100) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(101) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(102) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(103) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(104) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(105) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		// fset.File(106) : &token.File{name:"c", base:98, size:8, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}
		//
		// fset.File(107) : (*token.File)(nil)
	}

	for i := 0; i <= fset.Base(); i++ {
		pos := token.Pos(i)
		fmt.Printf("fset.Position(%03d) : %#v\n", pos, fset.Position(pos))
		// fset.Position(000) : token.Position{Filename:"", Offset:0, Line:0, Column:0}
		//
		// fset.Position(001) : token.Position{Filename:"a", Offset:0, Line:1, Column:1}
		// fset.Position(002) : token.Position{Filename:"a", Offset:1, Line:1, Column:2}
		// fset.Position(003) : token.Position{Filename:"a", Offset:2, Line:1, Column:3}
		// fset.Position(004) : token.Position{Filename:"a", Offset:3, Line:1, Column:4}
		// fset.Position(005) : token.Position{Filename:"a", Offset:4, Line:1, Column:5}
		// fset.Position(006) : token.Position{Filename:"a", Offset:5, Line:1, Column:6}
		// fset.Position(007) : token.Position{Filename:"a", Offset:6, Line:1, Column:7}
		// fset.Position(008) : token.Position{Filename:"a", Offset:7, Line:1, Column:8}
		// fset.Position(009) : token.Position{Filename:"a", Offset:8, Line:1, Column:9}
		// fset.Position(010) : token.Position{Filename:"a", Offset:9, Line:1, Column:10}
		// fset.Position(011) : token.Position{Filename:"a", Offset:10, Line:1, Column:11}
		// fset.Position(012) : token.Position{Filename:"a", Offset:11, Line:1, Column:12}
		// fset.Position(013) : token.Position{Filename:"a", Offset:12, Line:1, Column:13}
		// fset.Position(014) : token.Position{Filename:"a", Offset:13, Line:1, Column:14}
		// fset.Position(015) : token.Position{Filename:"a", Offset:14, Line:1, Column:15}
		// fset.Position(016) : token.Position{Filename:"a", Offset:15, Line:1, Column:16}
		// fset.Position(017) : token.Position{Filename:"a", Offset:16, Line:1, Column:17}
		// fset.Position(018) : token.Position{Filename:"a", Offset:17, Line:1, Column:18}
		// fset.Position(019) : token.Position{Filename:"a", Offset:18, Line:1, Column:19}
		// fset.Position(020) : token.Position{Filename:"a", Offset:19, Line:1, Column:20}
		// fset.Position(021) : token.Position{Filename:"a", Offset:20, Line:1, Column:21}
		// fset.Position(022) : token.Position{Filename:"a", Offset:21, Line:1, Column:22}
		// fset.Position(023) : token.Position{Filename:"a", Offset:22, Line:1, Column:23}
		// fset.Position(024) : token.Position{Filename:"a", Offset:23, Line:1, Column:24}
		// fset.Position(025) : token.Position{Filename:"a", Offset:24, Line:1, Column:25}
		// fset.Position(026) : token.Position{Filename:"a", Offset:25, Line:1, Column:26}
		// fset.Position(027) : token.Position{Filename:"a", Offset:26, Line:1, Column:27}
		// fset.Position(028) : token.Position{Filename:"a", Offset:27, Line:1, Column:28}
		// fset.Position(029) : token.Position{Filename:"a", Offset:28, Line:1, Column:29}
		// fset.Position(030) : token.Position{Filename:"a", Offset:29, Line:1, Column:30}
		// fset.Position(031) : token.Position{Filename:"a", Offset:30, Line:1, Column:31}
		// fset.Position(032) : token.Position{Filename:"a", Offset:31, Line:1, Column:32}
		// fset.Position(033) : token.Position{Filename:"a", Offset:32, Line:1, Column:33}
		// fset.Position(034) : token.Position{Filename:"a", Offset:33, Line:1, Column:34}
		// fset.Position(035) : token.Position{Filename:"a", Offset:34, Line:1, Column:35}
		// fset.Position(036) : token.Position{Filename:"a", Offset:35, Line:1, Column:36}
		// fset.Position(037) : token.Position{Filename:"a", Offset:36, Line:1, Column:37}
		// fset.Position(038) : token.Position{Filename:"a", Offset:37, Line:1, Column:38}
		// fset.Position(039) : token.Position{Filename:"a", Offset:38, Line:1, Column:39}
		// fset.Position(040) : token.Position{Filename:"a", Offset:39, Line:1, Column:40}
		// fset.Position(041) : token.Position{Filename:"a", Offset:40, Line:1, Column:41}
		// fset.Position(042) : token.Position{Filename:"a", Offset:41, Line:1, Column:42}
		// fset.Position(043) : token.Position{Filename:"a", Offset:42, Line:1, Column:43}
		// fset.Position(044) : token.Position{Filename:"a", Offset:43, Line:1, Column:44}
		// fset.Position(045) : token.Position{Filename:"a", Offset:44, Line:1, Column:45}
		// fset.Position(046) : token.Position{Filename:"a", Offset:45, Line:1, Column:46}
		// fset.Position(047) : token.Position{Filename:"a", Offset:46, Line:1, Column:47}
		// fset.Position(048) : token.Position{Filename:"a", Offset:47, Line:1, Column:48}
		// fset.Position(049) : token.Position{Filename:"a", Offset:48, Line:1, Column:49}
		// fset.Position(050) : token.Position{Filename:"a", Offset:49, Line:1, Column:50}
		// fset.Position(051) : token.Position{Filename:"a", Offset:50, Line:1, Column:51}
		// fset.Position(052) : token.Position{Filename:"a", Offset:51, Line:1, Column:52}
		// fset.Position(053) : token.Position{Filename:"a", Offset:52, Line:1, Column:53}
		// fset.Position(054) : token.Position{Filename:"a", Offset:53, Line:1, Column:54}
		// fset.Position(055) : token.Position{Filename:"a", Offset:54, Line:1, Column:55}
		// fset.Position(056) : token.Position{Filename:"a", Offset:55, Line:1, Column:56}
		// fset.Position(057) : token.Position{Filename:"a", Offset:56, Line:1, Column:57}
		// fset.Position(058) : token.Position{Filename:"a", Offset:57, Line:1, Column:58}
		// fset.Position(059) : token.Position{Filename:"a", Offset:58, Line:1, Column:59}
		// fset.Position(060) : token.Position{Filename:"a", Offset:59, Line:1, Column:60}
		// fset.Position(061) : token.Position{Filename:"a", Offset:60, Line:1, Column:61}
		// fset.Position(062) : token.Position{Filename:"a", Offset:61, Line:1, Column:62}
		// fset.Position(063) : token.Position{Filename:"a", Offset:62, Line:1, Column:63}
		// fset.Position(064) : token.Position{Filename:"a", Offset:63, Line:1, Column:64}
		// fset.Position(065) : token.Position{Filename:"a", Offset:64, Line:1, Column:65}
		// fset.Position(066) : token.Position{Filename:"a", Offset:65, Line:1, Column:66}
		// fset.Position(067) : token.Position{Filename:"a", Offset:66, Line:1, Column:67}
		// fset.Position(068) : token.Position{Filename:"a", Offset:67, Line:1, Column:68}
		//
		// fset.Position(069) : token.Position{Filename:"b", Offset:0, Line:1, Column:1}
		// fset.Position(070) : token.Position{Filename:"b", Offset:1, Line:1, Column:2}
		// fset.Position(071) : token.Position{Filename:"b", Offset:2, Line:1, Column:3}
		// fset.Position(072) : token.Position{Filename:"b", Offset:3, Line:1, Column:4}
		// fset.Position(073) : token.Position{Filename:"b", Offset:4, Line:1, Column:5}
		// fset.Position(074) : token.Position{Filename:"b", Offset:5, Line:1, Column:6}
		// fset.Position(075) : token.Position{Filename:"b", Offset:6, Line:1, Column:7}
		// fset.Position(076) : token.Position{Filename:"b", Offset:7, Line:1, Column:8}
		// fset.Position(077) : token.Position{Filename:"b", Offset:8, Line:1, Column:9}
		// fset.Position(078) : token.Position{Filename:"b", Offset:9, Line:1, Column:10}
		// fset.Position(079) : token.Position{Filename:"b", Offset:10, Line:1, Column:11}
		// fset.Position(080) : token.Position{Filename:"b", Offset:11, Line:1, Column:12}
		// fset.Position(081) : token.Position{Filename:"b", Offset:12, Line:1, Column:13}
		// fset.Position(082) : token.Position{Filename:"b", Offset:13, Line:1, Column:14}
		// fset.Position(083) : token.Position{Filename:"b", Offset:14, Line:1, Column:15}
		// fset.Position(084) : token.Position{Filename:"b", Offset:15, Line:1, Column:16}
		// fset.Position(085) : token.Position{Filename:"b", Offset:16, Line:1, Column:17}
		// fset.Position(086) : token.Position{Filename:"b", Offset:17, Line:1, Column:18}
		// fset.Position(087) : token.Position{Filename:"b", Offset:18, Line:1, Column:19}
		// fset.Position(088) : token.Position{Filename:"b", Offset:19, Line:1, Column:20}
		// fset.Position(089) : token.Position{Filename:"b", Offset:20, Line:1, Column:21}
		// fset.Position(090) : token.Position{Filename:"b", Offset:21, Line:1, Column:22}
		// fset.Position(091) : token.Position{Filename:"b", Offset:22, Line:1, Column:23}
		// fset.Position(092) : token.Position{Filename:"b", Offset:23, Line:1, Column:24}
		// fset.Position(093) : token.Position{Filename:"b", Offset:24, Line:1, Column:25}
		// fset.Position(094) : token.Position{Filename:"b", Offset:25, Line:1, Column:26}
		// fset.Position(095) : token.Position{Filename:"b", Offset:26, Line:1, Column:27}
		// fset.Position(096) : token.Position{Filename:"b", Offset:27, Line:1, Column:28}
		// fset.Position(097) : token.Position{Filename:"b", Offset:28, Line:1, Column:29}
		//
		// fset.Position(098) : token.Position{Filename:"c", Offset:0, Line:1, Column:1}
		// fset.Position(099) : token.Position{Filename:"c", Offset:1, Line:1, Column:2}
		// fset.Position(100) : token.Position{Filename:"c", Offset:2, Line:1, Column:3}
		// fset.Position(101) : token.Position{Filename:"c", Offset:3, Line:1, Column:4}
		// fset.Position(102) : token.Position{Filename:"c", Offset:4, Line:1, Column:5}
		// fset.Position(103) : token.Position{Filename:"c", Offset:5, Line:1, Column:6}
		// fset.Position(104) : token.Position{Filename:"c", Offset:6, Line:1, Column:7}
		// fset.Position(105) : token.Position{Filename:"c", Offset:7, Line:1, Column:8}
		// fset.Position(106) : token.Position{Filename:"c", Offset:8, Line:1, Column:9}
		//
		// fset.Position(107) : token.Position{Filename:"", Offset:0, Line:0, Column:0}
	}

	fset.Iterate(func(f *token.File) bool {
		fmt.Printf("\nf : %#v\n\n", f)
		// f : &token.File{name:"a", base:0, size:11, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}

		return f.Name() == "b"
	})
}
