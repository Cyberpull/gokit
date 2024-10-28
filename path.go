package gokit

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type PathFn func() (path string, err error)

type xPath struct {
	//
}

func (x xPath) Delim() string {
	return string([]rune{os.PathSeparator})
}

func (x xPath) Join(paths ...any) string {
	delim := x.Delim()

	return JoinFunc(delim, paths, func(v string) string {
		v = strings.TrimPrefix(v, delim)
		v = strings.TrimSuffix(v, delim)
		return v
	})
}

func (x xPath) JoinPrefixed(paths ...any) string {
	result := x.Join(paths...)
	return x.Delim() + result
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
	path = x.sanitize(path)

	return
}

// ================================

var Path xPath
