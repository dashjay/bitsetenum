package bitsetenum

import (
	"errors"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"strings"
)

type Style int

const (
	StyleSetAndUnset Style = 0 + iota
	StyleSetWithValue
)

type EnumBitsetOptions struct {
	TypeName        string
	EnumNames       []string
	StepIn          int
	EnumBitsetStyle Style
}

func checkEBOValidAndSetDefault(e *EnumBitsetOptions) error {
	if e.TypeName == "" {
		return errors.New("TypeName is needed")
	}
	if len(e.EnumNames) == 0 {
		return errors.New("EnumNames is needed")
	}
	if e.StepIn == 0 {
		e.StepIn = 1
	}
	return nil
}

type OptionFunc func(e *EnumBitsetOptions)

func WithTypeName(TypeName string) OptionFunc {
	return func(e *EnumBitsetOptions) {
		e.TypeName = TypeName
	}
}

func WithEnumNames(EnumNames ...string) OptionFunc {
	return func(e *EnumBitsetOptions) {
		e.EnumNames = EnumNames
	}
}

func WithStepIn(StepIn int) OptionFunc {
	return func(e *EnumBitsetOptions) {
		e.StepIn = StepIn
	}
}
func WithEnumBitsetStyle(style Style) OptionFunc {
	return func(e *EnumBitsetOptions) {
		e.EnumBitsetStyle = style
	}
}

func RenderEnumBitset(r *Render, opts ...OptionFunc) error {
	var ebo EnumBitsetOptions
	for _, opt := range opts {
		opt(&ebo)
	}
	err := checkEBOValidAndSetDefault(&ebo)
	if err != nil {
		return err
	}
	r.File().Type().Id(ebo.TypeName).Uint64()
	r.File().Const().Defs(
		func(enums []string) []Code {
			stepIn := 0
			var out []Code
			for i := 0; i < len(enums); i++ {
				out = append(out, Id(enums[i]).
					Id(ebo.TypeName).
					Op("=").
					Lit(1).
					Op("<<").
					Lit(stepIn))
				stepIn += ebo.StepIn
			}
			return out
		}(ebo.EnumNames)...,
	)
	receiver := strings.ToLower(string(ebo.TypeName[0]))

	for i := 0; i < len(ebo.EnumNames); i++ {
		if ebo.EnumBitsetStyle == StyleSetAndUnset {
			r.File().Func().Params(Id(receiver).Op("*").Id(ebo.TypeName)).Id(fmt.Sprintf("Set%s%s", ebo.TypeName, ebo.EnumNames[i])).Params().Block(Id("*").Id(receiver).Op("|=").Id(ebo.EnumNames[i]))
			r.File().Func().Params(Id(receiver).Op("*").Id(ebo.TypeName)).Id(fmt.Sprintf("UnSet%s%s", ebo.TypeName, ebo.EnumNames[i])).Params().Block(Id("*").Id(receiver).Op("&=").Op("^").Id(ebo.EnumNames[i]))
		} else {
			r.File().Func().Params(Id(receiver).Id("*").Id(ebo.TypeName)).
				Id(fmt.Sprintf("Set%s%s", ebo.TypeName, ebo.EnumNames[i])).
				Params(Id("val").Bool()).
				Block(
					If(Id("val")).Block(Id("*").Id(receiver).Op("|=").Id(ebo.EnumNames[i])).Else().Block(Id("*").Id(receiver).Op("&=").Op("^").Id(ebo.EnumNames[i])),
				)
		}
		r.File().Func().Params(Id(receiver).Id("*").Id(ebo.TypeName)).Id(fmt.Sprintf("Get%s%s", ebo.TypeName, ebo.EnumNames[i])).Params().Bool().Block(Return().Id(fmt.Sprintf("(*%s)", receiver)).Id("&").Id(ebo.EnumNames[i]).Id("!=").Lit(0))
	}
	return nil
}
