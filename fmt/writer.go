package fmt

import (
	"io"

	"github.com/Cyberpull/gokit/color"
)

type xColorFmt struct {
	color color.Color
}

func (x *xColorFmt) Fprint(w io.Writer, a ...any) (n int, err error) {
	return x.color.Fprint(w, a...)
}

func (x *xColorFmt) Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	return x.color.Fprintf(w, format, a...)
}

func (x *xColorFmt) Fprintln(w io.Writer, a ...any) (n int, err error) {
	return x.color.Fprintln(w, a...)
}

func (x *xColorFmt) Print(a ...any) (n int, err error) {
	return x.color.Print(a...)
}

func (x *xColorFmt) Printf(format string, a ...any) (n int, err error) {
	return x.color.Printf(format, a...)
}

func (x *xColorFmt) Println(a ...any) (n int, err error) {
	return x.color.Println(a...)
}

func (x *xColorFmt) Sprint(a ...any) string {
	return x.color.Sprint(a...)
}

func (x *xColorFmt) Sprintf(format string, a ...any) string {
	return x.color.Sprintf(format, a...)
}

func (x *xColorFmt) Sprintln(a ...any) string {
	return x.color.Sprintln(a...)
}

// =====================================

func Color(a ...color.Attribute) *xColorFmt {
	return &xColorFmt{color: *color.New(a...)}
}
