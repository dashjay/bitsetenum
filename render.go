package bitsetenum

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"os"
)

type Render struct {
	f *File
}

func (r *Render) File() *File {
	return r.f
}

func NewRenderWithPackageName(pkgName string) *Render {
	return &Render{f: NewFile(pkgName)}
}

func (r *Render) RenderToFile(fp string) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "%#v", r.f)
	return err
}
