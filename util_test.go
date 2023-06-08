package wopan_test

import (
	"testing"

	"github.com/Xhofe/wopan-sdk-go"
)

func TestCopyStruct(t *testing.T) {
	type A struct {
		Name string
		Age  int
		A    string
	}
	type B struct {
		Name string
		Age  int
		B    string
	}
	a := A{
		Name: "test",
		Age:  18,
	}
	b := B{}
	wopan.CopyStruct(a, &b)
	if b.Name != "test" || b.Age != 18 {
		t.Error("copy struct failed")
	}
}
