package gokit

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type PathFn func() (path string, err error)

type xPath struct {
	//
}

func (x xPath) Join(paths ...any) string {
	delim := string([]rune{os.PathSeparator})

	var buff bytes.Buffer

	for _, path := range paths {
		data := fmt.Sprint(path)
		data = strings.TrimSpace(data)

		if data == "" {
			continue
		}

		if buff.Len() > 0 {
			buff.WriteString(delim)
		}

		buff.WriteString(data)
	}

	return buff.String()
}

func (x xPath) FromExecutable(paths ...any) (file string, err error) {
	return x.get(
		x.getExecPathFromArgs(paths...),
		x.getExecPathFromCaller(0, paths...),
		x.getExecPathFromCaller(1, paths...),
		x.getExecPathFromSource(paths...),
	)
}

func (x xPath) IsDir(file string) bool {
	info, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func (x xPath) IsFile(file string) bool {
	info, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func (x xPath) get(fns ...PathFn) (path string, err error) {
	for _, fn := range fns {
		path, err = fn()

		if err != nil {
			continue
		}

		if x.IsFile(path) {
			break
		}
	}

	return
}

func (x xPath) getExecPathFromCaller(skip int, paths ...any) PathFn {
	return func() (path string, err error) {
		_, file, _, _ := runtime.Caller(skip)
		return x.getAbsolutePathFromFile(file, paths...)
	}
}

func (x xPath) getExecPathFromArgs(paths ...any) PathFn {
	return func() (path string, err error) {
		return x.getAbsolutePathFromFile(os.Args[0], paths...)
	}
}

func (x xPath) getExecPathFromSource(paths ...any) PathFn {
	return func() (path string, err error) {
		return x.getAbsolutePathFromDir("./", paths...)
	}
}

func (x xPath) getAbsolutePathFromFile(fromFile string, paths ...any) (path string, err error) {
	dir := filepath.Dir(fromFile)
	return x.getAbsolutePathFromDir(dir, paths...)
}

func (x xPath) getAbsolutePathFromDir(dir string, paths ...any) (path string, err error) {
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}

	allPaths := append([]any{dir}, paths...)

	path = x.Join(allPaths...)

	return
}

// ================================

var Path xPath
