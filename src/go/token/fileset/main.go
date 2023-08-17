package main

import (
	"fmt"
	"go/token"
)

func main() {
	fSet := token.FileSet{}
	fmt.Printf("fSet : %#v\n", fSet)
	// fSet : token.FileSet{mutex:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:atomic.Int32{_:atomic.noCopy{}, v:0}, readerWait:atomic.Int32{_:atomic.noCopy{}, v:0}}, base:0, files:[]*token.File(nil), last:atomic.Pointer[go/token.File]{_:[0]*token.File{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(nil)}}

	fSet.AddFile("a", -1, 1)
	fmt.Printf("fSet : %#v\n", fSet)
	// fSet : token.FileSet{mutex:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:atomic.Int32{_:atomic.noCopy{}, v:0}, readerWait:atomic.Int32{_:atomic.noCopy{}, v:0}}, base:2, files:[]*token.File{(*token.File)(0xc00009c120)}, last:atomic.Pointer[go/token.File]{_:[0]*token.File{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(0xc00009c120)}}

	fSet.AddFile("b", -1, 10)
	fmt.Printf("fSet : %#v\n", fSet)
	// fSet : token.FileSet{mutex:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:atomic.Int32{_:atomic.noCopy{}, v:0}, readerWait:atomic.Int32{_:atomic.noCopy{}, v:0}}, base:13, files:[]*token.File{(*token.File)(0xc00009c120), (*token.File)(0xc00009c180)}, last:atomic.Pointer[go/token.File]{_:[0]*token.File{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(0xc00009c180)}}

	fSet.AddFile("c", -1, 1)
	fmt.Printf("fSet : %#v\n", fSet)
	// fSet : token.FileSet{mutex:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:atomic.Int32{_:atomic.noCopy{}, v:0}, readerWait:atomic.Int32{_:atomic.noCopy{}, v:0}}, base:15, files:[]*token.File{(*token.File)(0xc00009c120), (*token.File)(0xc00009c180), (*token.File)(0xc00009c1e0)}, last:atomic.Pointer[go/token.File]{_:[0]*token.File{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(0xc00009c1e0)}}

	fmt.Printf("file.File(1) : %#v\n", fSet.File(1))
	// file.File(1) : &token.File{name:"a", base:0, size:1, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}

	fmt.Printf("file.File(2) : %#v\n", fSet.File(2))
	// file.File(2) : &token.File{name:"b", base:2, size:10, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}

	fmt.Printf("file.File(13) : %#v\n", fSet.File(13))
	// file.File(13) : &token.File{name:"c", base:13, size:1, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}

	fmt.Printf("file.File(14) : %#v\n", fSet.File(14))

	fSet.Iterate(func(f *token.File) bool {
		fmt.Printf("\nf : %#v\n\n", f)
		// f : &token.File{name:"a", base:0, size:1, mutex:sync.Mutex{state:0, sema:0x0}, lines:[]int{0}, infos:[]token.lineInfo(nil)}

		return f.Name() == "b"
	})

	fmt.Printf("fSet.Position(0) : %#v\n", fSet.Position(0))
	// fSet.Position(0) : token.Position{Filename:"", Offset:0, Line:0, Column:0}

	fmt.Printf("fSet.Position(1) : %#v\n", fSet.Position(1))
	// fSet.Position(1) : token.Position{Filename:"a", Offset:1, Line:1, Column:2}

	fmt.Printf("fSet.Position(2) : %#v\n", fSet.Position(2))
	// fSet.Position(2) : token.Position{Filename:"b", Offset:0, Line:1, Column:1}

	fmt.Printf("fSet.Position(3) : %#v\n", fSet.Position(3))
	// fSet.Position(3) : token.Position{Filename:"b", Offset:1, Line:1, Column:2}

	fmt.Printf("fSet.Position(4) : %#v\n", fSet.Position(4))
	// fSet.Position(4) : token.Position{Filename:"b", Offset:2, Line:1, Column:3}

	fmt.Printf("fSet.Position(5) : %#v\n", fSet.Position(5))
	// fSet.Position(5) : token.Position{Filename:"b", Offset:3, Line:1, Column:4}

	fmt.Printf("fSet.Position(6) : %#v\n", fSet.Position(6))
	// fSet.Position(6) : token.Position{Filename:"b", Offset:4, Line:1, Column:5}

	fmt.Printf("fSet.Position(7) : %#v\n", fSet.Position(7))
	// fSet.Position(7) : token.Position{Filename:"b", Offset:5, Line:1, Column:6}

	fmt.Printf("fSet.Position(8) : %#v\n", fSet.Position(8))
	// fSet.Position(8) : token.Position{Filename:"b", Offset:6, Line:1, Column:7}

	fmt.Printf("fSet.Position(9) : %#v\n", fSet.Position(9))
	// fSet.Position(9) : token.Position{Filename:"b", Offset:7, Line:1, Column:8}

	fmt.Printf("fSet.Position(10) : %#v\n", fSet.Position(10))
	// fSet.Position(10) : token.Position{Filename:"b", Offset:8, Line:1, Column:9}

	fmt.Printf("fSet.Position(11) : %#v\n", fSet.Position(11))
	// fSet.Position(11) : token.Position{Filename:"b", Offset:9, Line:1, Column:10}

	fmt.Printf("fSet.Position(12) : %#v\n", fSet.Position(12))
	// fSet.Position(12) : token.Position{Filename:"b", Offset:10, Line:1, Column:11}

	fmt.Printf("fSet.Position(13) : %#v\n", fSet.Position(13))
	// fSet.Position(13) : token.Position{Filename:"c", Offset:0, Line:1, Column:1}

	fmt.Printf("fSet.Position(14) : %#v\n", fSet.Position(14))
	// fSet.Position(14) : token.Position{Filename:"c", Offset:1, Line:1, Column:2}

	fmt.Printf("fSet.Position(15) : %#v\n", fSet.Position(15))
	// fSet.Position(15) : token.Position{Filename:"", Offset:0, Line:0, Column:0}

	fmt.Printf("fSet.Position(16) : %#v\n", fSet.Position(16))
	// fSet.Position(16) : token.Position{Filename:"", Offset:0, Line:0, Column:0}

	fmt.Printf("fSet.Position(17) : %#v\n", fSet.Position(17))
	// fSet.Position(17) : token.Position{Filename:"", Offset:0, Line:0, Column:0}

	fmt.Printf("fSet.Position(18) : %#v\n", fSet.Position(18))
	// fSet.Position(18) : token.Position{Filename:"", Offset:0, Line:0, Column:0}
}
