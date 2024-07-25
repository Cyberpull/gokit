//go:build !windows

package gokit

import (
	"os"
	"os/user"
	"strings"
)

func (x xPath) Expand(path string) string {
	if path == "" {
		return path
	}

	u, _ := user.Current()

	if u == nil || u.HomeDir == "" {
		return path
	}

	var home string

	if char := rune(path[0]); char == '~' {
		delim := string([]rune{os.PathSeparator})
		home = strings.TrimRight(u.HomeDir, delim)
		path = path[1:]
	}

	return home + os.ExpandEnv(path)
}
