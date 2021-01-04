package objmodelexp

import (
    "fmt"
    "reflect"
)

func PrintObjectDumpTableHeader() {
    fmt.Printf("%-12s%-15s%-20s%-10s %-11s %-4s\n", "Var", "Type", "Address", "RootOffset", "LocalOffset", "Size")
}

func DumpObject(name string, p reflect.Value) {
    v := p.Elem()
    dumpObject(name, v, v.UnsafeAddr(), v.UnsafeAddr())
}

func dumpObject(path string, v reflect.Value, rootBaseAddr uintptr, localBaseAddr uintptr) {
    dumpObjectDetail(path, v, rootBaseAddr, localBaseAddr)

    switch v.Kind() {
    case reflect.Struct:
        childLocalBaseAddr := v.UnsafeAddr()

        for i := 0; i < v.NumField(); i++ {
            fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
            dumpObject(fieldPath, v.Field(i), rootBaseAddr, childLocalBaseAddr)
        }
    }
}

func dumpObjectDetail(path string, v reflect.Value, rootBaseAddr uintptr, localBaseAddr uintptr) {
    fmt.Printf("%-12s%-15s0x%016x  %10x %11x %4v\n", path, v.Type().String(), v.UnsafeAddr(),
        v.UnsafeAddr() - rootBaseAddr, v.UnsafeAddr() - localBaseAddr, v.Type().Size())
}