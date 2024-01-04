package main

import (
	"github.com/dashjay/bitsetenum"
)

//go:generate go run main.go
func main() {
	rd := bitsetenum.NewRenderWithPackageName("main_test")
	err := bitsetenum.RenderEnumBitset(rd,
		bitsetenum.WithTypeName("Status"),
		bitsetenum.WithEnumNames("Creating", "Normal", "Updating", "CanaryUpdating"),
		bitsetenum.WithEnumBitsetStyle(bitsetenum.StyleSetWithValue),
	)
	if err != nil {
		panic(err)
	}
	err = rd.RenderToFile("main_test/all_test.go")
	if err != nil {
		panic(err)
	}
}
