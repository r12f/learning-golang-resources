package objmodelexp

import (
    "fmt"
    "reflect"
)

func DumpObjectModelSimpleDataType() {
    PrintObjectDumpTableHeader()

    var b bool
    DumpObject("b", reflect.ValueOf(&b))

    var i int
    DumpObject("i", reflect.ValueOf(&i))

    var i8 int8
    DumpObject("i8", reflect.ValueOf(&i8))

    var i16 int16
    DumpObject("i16", reflect.ValueOf(&i16))

    var i32 int32
    DumpObject("i32", reflect.ValueOf(&i32))

    var i64 int64
    DumpObject("i64", reflect.ValueOf(&i64))

    var u8 uint8
    DumpObject("u8", reflect.ValueOf(&u8))

    var u16 uint16
    DumpObject("u16", reflect.ValueOf(&u16))

    var u32 uint32
    DumpObject("u32", reflect.ValueOf(&u32))

    var u64 uint64
    DumpObject("u64", reflect.ValueOf(&u64))

    var p uintptr
    DumpObject("p", reflect.ValueOf(&p))

    var f32 float32
    DumpObject("f32", reflect.ValueOf(&f32))

    var f64 float64
    DumpObject("f64", reflect.ValueOf(&f64))

    var c64 complex64
    DumpObject("c64", reflect.ValueOf(&c64))

    var c128 complex128
    DumpObject("c128", reflect.ValueOf(&c128))

    var s string = "SomeRandomString"
    DumpObject("s", reflect.ValueOf(&s))

    var r rune
    DumpObject("r", reflect.ValueOf(&r))
}

func DumpObjectModelCompositeDataType() {
    PrintObjectDumpTableHeader()

    var a0 [0]int
    DumpObject("a0", reflect.ValueOf(&a0))

    var a1 [1]int
    DumpObject("a1", reflect.ValueOf(&a1))

    var a5 [5]int
    DumpObject("a5", reflect.ValueOf(&a5))

    var b5 [5]bool
    DumpObject("b5", reflect.ValueOf(&b5))

    var si []int
    DumpObject("si", reflect.ValueOf(&si))

    var si32 []int32
    DumpObject("si32", reflect.ValueOf(&si32))

    var mii map[int]int = make(map[int]int)
    DumpObject("mii", reflect.ValueOf(&mii))

    var mi32i map[int32]int = make(map[int32]int)
    DumpObject("mi32i", reflect.ValueOf(&mi32i))

    var mii32 map[int]int32 = make(map[int]int32)
    DumpObject("mii32", reflect.ValueOf(&mii32))

    var ci chan int = make(chan int)
    DumpObject("ci", reflect.ValueOf(&ci))

    var cb chan bool = make(chan bool)
    DumpObject("cb", reflect.ValueOf(&cb))

    var f func() = CheckFunctionInternal
    DumpObject("f", reflect.ValueOf(&f))

    var i int = 1
    var cls func() = func() {
        i = 2
        fmt.Printf("i = %v\n", i)
    }
    DumpObject("cls", reflect.ValueOf(&cls))
}

func CheckSliceInternal() {
    var s []int
    UpdateSlice(s)
    fmt.Printf("AfterReturn: %v (Ptr = %p)\n", s, &s)
}

//go:noinline
func UpdateSlice(s []int) {
    s = append(s, 1)
    fmt.Printf("AfterUpdate: %v (Ptr = %p)\n", s, &s)
}

func CheckMapInternal() {
    // var m map[int]int // This will cause a runtime error
    var m map[int]int = make(map[int]int)
    m[5] = 3
    fmt.Printf("BeforeCall: %v (Ptr = %p)\n", m, &m)
    UpdateMap(m)
    fmt.Printf("AfterReturn: %v (Ptr = %p)\n", m, &m)
}

//go:noinline
func UpdateMap(m map[int]int) {
    m = make(map[int]int)
    fmt.Printf("AfterUpdate: %v (Ptr = %p)\n", m, &m)
}

func CheckChannelInternal() {
    var c chan int = make(chan int, 1)
    UpdateChannel(c)
    fmt.Printf("AfterReturn: Ptr = %p\n", &c)
}

//go:noinline
func UpdateChannel(c chan int) {
    c = make(chan int, 1)
    fmt.Printf("AfterUpdate: Ptr = %p\n", &c)
}

func CheckFunctionInternal() {
    var f func() = CheckFunctionInternal
    f()
    UpdateFunction(f)
    fmt.Printf("AfterReturn: Ptr = %p\n", &f)
}

//go:noinline
func UpdateFunction(f func()) {
    f = CheckChannelInternal
    fmt.Printf("AfterUpdate: Ptr = %p\n", &f)
}

func CheckClosureInternal() {
    var i int = 1
    var f func() = func() {
        i = 2
        fmt.Printf("i = %v\n", i)
    }
    UpdateClosure(f)
    fmt.Printf("AfterReturn: Ptr = %p\n", &f)
}

//go:noinline
func UpdateClosure(f func()) {
    var b bool = false
    f = func() {
        b = true
        fmt.Printf("i = %v\n", b)
    }
    fmt.Printf("AfterUpdate: Ptr = %p\n", &f)
}

func DumpObjectModelSingleObject() {
    PrintObjectDumpTableHeader()
    var a A
    DumpObject("a", reflect.ValueOf(&a))
    fmt.Println()

    PrintObjectDumpTableHeader()
    var b B
    DumpObject("b", reflect.ValueOf(&b))
}

func DumpObjectModelSingleEmbedding() {
    PrintObjectDumpTableHeader()

    var c C
    DumpObject("c", reflect.ValueOf(&c))
}

func DumpObjectModelSingleNestedEmbedding() {
    PrintObjectDumpTableHeader()

    var d D
    DumpObject("d", reflect.ValueOf(&d))
}
