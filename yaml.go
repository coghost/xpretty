package xpretty

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/printer"
	"github.com/mattn/go-colorable"
	"github.com/tidwall/gjson"
)

const _escape = "\x1b"

type PrettyOpts struct {
	withLineNumber bool
	// json indent
	indent int
	// filter out json keys
	filterKey string
}

type PrettyOptFunc func(o *PrettyOpts)

func bindPrettyOpts(opt *PrettyOpts, opts ...PrettyOptFunc) {
	for _, f := range opts {
		f(opt)
	}
}

func WithLineNumber(b bool) PrettyOptFunc {
	return func(o *PrettyOpts) {
		o.withLineNumber = b
	}
}

func WithFilterKey(s string) PrettyOptFunc {
	return func(o *PrettyOpts) {
		o.filterKey = s
	}
}

func WithIndent(i int) PrettyOptFunc {
	return func(o *PrettyOpts) {
		o.indent = i
	}
}

func _format(attr color.Attribute) string {
	return fmt.Sprintf("%s[%dm", _escape, attr)
}

func PrettyYaml(rawYaml string, opts ...PrettyOptFunc) error {
	opt := PrettyOpts{withLineNumber: false}
	bindPrettyOpts(&opt, opts...)

	tokens := lexer.Tokenize(rawYaml)
	var p printer.Printer
	p.LineNumber = opt.withLineNumber
	p.LineNumberFormat = func(num int) string {
		fn := color.New(color.Bold, color.Faint).SprintFunc()
		return fn(fmt.Sprintf("%2d | ", num))
	}
	p.Bool = func() *printer.Property {
		return &printer.Property{
			Prefix: _format(color.FgHiMagenta),
			Suffix: _format(color.Reset),
		}
	}
	p.Number = func() *printer.Property {
		return &printer.Property{
			Prefix: _format(color.FgHiMagenta),
			Suffix: _format(color.Reset),
		}
	}
	p.MapKey = func() *printer.Property {
		return &printer.Property{
			Prefix: _format(color.FgHiCyan),
			Suffix: _format(color.Reset),
		}
	}
	p.Anchor = func() *printer.Property {
		return &printer.Property{
			Prefix: _format(color.FgHiYellow),
			Suffix: _format(color.Reset),
		}
	}
	p.Alias = func() *printer.Property {
		return &printer.Property{
			Prefix: _format(color.FgHiYellow),
			Suffix: _format(color.Reset),
		}
	}
	p.String = func() *printer.Property {
		return &printer.Property{
			Prefix: _format(color.FgHiGreen),
			Suffix: _format(color.Reset),
		}
	}
	writer := colorable.NewColorableStdout()
	writer.Write([]byte(p.PrintTokens(tokens) + "\n"))
	return nil
}

func MustPrettyYaml(raw string, opts ...PrettyOptFunc) {
	e := PrettyYaml(raw, opts...)
	if e != nil {
		fmt.Printf("%v\n", yaml.FormatError(e, true, true))
		panic(e)
	}
}

func JsonToYaml(raw string) (string, error) {
	v, e := yaml.JSONToYAML([]byte(raw))
	return string(v), e
}

func MustJsonToYaml(raw string) string {
	v, e := yaml.JSONToYAML([]byte(raw))
	if e != nil {
		panic(e)
	}
	return string(v)
}

func PrettyJsonAsYaml(rawJson string, opts ...PrettyOptFunc) {
	opt := PrettyOpts{filterKey: ""}
	bindPrettyOpts(&opt, opts...)

	raw := rawJson
	if opt.filterKey != "" {
		raw = gjson.Get(rawJson, opt.filterKey).String()
	}
	v := MustJsonToYaml(raw)
	MustPrettyYaml(string(v), opts...)
}
