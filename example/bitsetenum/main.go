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
		bitsetenum.WithAssertion(),
	)

	err = bitsetenum.RenderEnumBitset(rd,
		bitsetenum.WithTypeName("StatusWithUnset"),
		bitsetenum.WithEnumNames("Creating1", "Normal1", "Updating1", "CanaryUpdating1"),
		bitsetenum.WithEnumBitsetStyle(bitsetenum.StyleSetAndUnset),
		bitsetenum.WithAssertion(),
	)
	if err != nil {
		panic(err)
	}
	err = rd.RenderToFile("main_test/all_test.go")
	if err != nil {
		panic(err)
	}
}
