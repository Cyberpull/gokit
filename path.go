package gokit

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type PathFn func() (path string, err error)

func Path(paths ...string) string {
	delim := string([]rune{os.PathSeparator})
	return strings.Join(paths, delim)
}

func PathFromExecutable(paths ...string) (file string, err error) {
	return getPath(
		getExecPathFromArgs(paths...),
		getExecPathFromCaller(0, paths...),
		getExecPathFromCaller(1, paths...),
		getExecPathFromSource(paths...),
	)
}

func IsDir(file string) bool {
	info, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func IsFile(file string) bool {
	info, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func getPath(fns ...PathFn) (path string, err error) {
	for _, fn := range fns {
		path, err = fn()

		if err != nil {
			continue
		}

		if IsFile(path) {
			break
		}
	}

	return
}

func getExecPathFromCaller(skip int, paths ...string) PathFn {
	return func() (path string, err error) {
		_, file, _, _ := runtime.Caller(skip)
		return getAbsolutePathFromFile(file, paths...)
	}
}

func getExecPathFromArgs(paths ...string) PathFn {
	return func() (path string, err error) {
		return getAbsolutePathFromFile(os.Args[0], paths...)
	}
}

func getExecPathFromSource(paths ...string) PathFn {
	return func() (path string, err error) {
		return getAbsolutePathFromDir("./", paths...)
	}
}

func getAbsolutePathFromFile(fromFile string, paths ...string) (path string, err error) {
	dir := filepath.Dir(fromFile)
	return getAbsolutePathFromDir(dir, paths...)
}

func getAbsolutePathFromDir(dir string, paths ...string) (path string, err error) {
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}

	allPaths := append([]string{dir}, paths...)

	path = Path(allPaths...)

	return
}
