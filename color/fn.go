package color

import "github.com/fatih/color"

func String(a Attribute, v ...any) string {
	c := color.New(color.Attribute(a))
	return c.SprintFunc()(v...)
}

func Stringf(a Attribute, format string, v ...any) string {
	c := color.New(color.Attribute(a))
	return c.SprintfFunc()(format, v...)
}

func a(v ...Attribute) []color.Attribute {
	value := make([]color.Attribute, 0)

	for _, a := range v {
		value = append(value, color.Attribute(a))
	}

	return value
}
