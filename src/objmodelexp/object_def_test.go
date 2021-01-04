package objmodelexp

import (
    "testing"
)

func TestObjectModelSimpleDataType(t *testing.T) {
    DumpObjectModelSimpleDataType()
}

func TestObjectModelCompositeDataType(t *testing.T) {
    DumpObjectModelCompositeDataType()
}

func TestObjectModelFunctionCallWithSlice(t *testing.T) {
    CheckSliceInternal()
}

func TestObjectModelFunctionCallWithMap(t *testing.T) {
    CheckMapInternal()
}

func TestObjectModelFunctionCallWithChannel(t *testing.T) {
    CheckChannelInternal()
}

func TestObjectModelFunctionCallWithFunction(t *testing.T) {
    CheckFunctionInternal()
}

func TestObjectModelClosureCallWithClosure(t *testing.T) {
    CheckClosureInternal()
}

func TestObjectModelSingleObject(t *testing.T) {
    DumpObjectModelSingleObject()
}

func TestObjectModelSingleEmbedding(t *testing.T) {
    DumpObjectModelSingleEmbedding()
}

func TestObjectModelSingleNestedEmbedding(t *testing.T) {
    DumpObjectModelSingleNestedEmbedding()
}
