package gokit

import (
	"bytes"
	"os"
	"unicode"
)

func (x xPath) Expand(path string) string {
	if path == "" {
		return path
	}

	var isEnv bool
	var buff, env bytes.Buffer

	for _, char := range path {
		if char == '%' {
			if !isEnv {
				isEnv = true
				continue
			}

			isEnv = false

			value := os.Getenv(env.String())
			buff.WriteString(value)

			env.Reset()
			continue
		}

		if isEnv {
			char = unicode.ToUpper(char)
			env.WriteRune(char)
			continue
		}

		buff.WriteRune(char)
	}

	return buff.String()
}

func (x xPath) sanitize(path string) string {
	return path
}
