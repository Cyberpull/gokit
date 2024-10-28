package log

import (
	"io"

	"github.com/Cyberpull/gokit/color"
)

type Options struct {
	out    io.Writer
	prefix string
	flag   int
	color  color.Attribute
}
