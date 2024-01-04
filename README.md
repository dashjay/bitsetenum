# bitsetenum

A useful tool for generating codes about bitset and enum in golang

# Usage

like example, you can write code like this and then run go generate ./...

```go

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
```

bitsetenum will help you to generate file like

```go
package main_test

type Status uint64

const (
	Creating       Status = 1 << 0
	Normal         Status = 1 << 1
	Updating       Status = 1 << 2
	CanaryUpdating Status = 1 << 3
)

func (s *Status) SetStatusCreating(val bool) {
	if val {
		*s |= Creating
	} else {
		*s &= ^Creating
	}
}
func (s *Status) GetStatusCreating() bool {
	return (*s)&Creating != 0
}
func (s *Status) SetStatusNormal(val bool) {
	if val {
		*s |= Normal
	} else {
		*s &= ^Normal
	}
}
func (s *Status) GetStatusNormal() bool {
	return (*s)&Normal != 0
}
func (s *Status) SetStatusUpdating(val bool) {
	if val {
		*s |= Updating
	} else {
		*s &= ^Updating
	}
}
func (s *Status) GetStatusUpdating() bool {
	return (*s)&Updating != 0
}
func (s *Status) SetStatusCanaryUpdating(val bool) {
	if val {
		*s |= CanaryUpdating
	} else {
		*s &= ^CanaryUpdating
	}
}
func (s *Status) GetStatusCanaryUpdating() bool {
	return (*s)&CanaryUpdating != 0
}
```

this can be useful when we want a bitset and enum but don't want to copy a lot of codes
then you can easily use it anywhere in your codes


