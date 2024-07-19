package gokit

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type CYBReaderLike interface {
	ReadBytes(delim byte) ([]byte, error)
	ReadLine() (b []byte, err error)
	ReadString(delim byte) (s string, err error)
	ReadStringLine() (s string, err error)
}

type IOData[T any] struct {
	Error error
	Data  T
}

type xIO struct {
	//
}

func (x xIO) ReadBytes(r io.Reader, delim byte) (value chan IOData[[]byte]) {
	value = make(chan IOData[[]byte], 1)

	go func() {
		var resp IOData[[]byte]

		switch reader := r.(type) {
		case CYBReaderLike:
			resp.Data, resp.Error = reader.ReadBytes(delim)

		default:
			buff := bufio.NewReader(r)
			resp.Data, resp.Error = buff.ReadBytes(delim)

			if resp.Error == nil {
				resp.Data = bytes.TrimSuffix(resp.Data, []byte{delim})
			}
		}

		value <- resp
	}()

	return
}

func (x xIO) ReadLine(r io.Reader) (value chan IOData[[]byte]) {
	value = make(chan IOData[[]byte], 1)

	go func() {
		var resp IOData[[]byte]

		switch reader := r.(type) {
		case CYBReaderLike:
			resp.Data, resp.Error = reader.ReadLine()

		default:
			buff := bufio.NewReader(r)
			resp.Data, resp.Error = buff.ReadBytes('\n')

			if resp.Error == nil {
				resp.Data = bytes.TrimSuffix(resp.Data, []byte{'\n'})
			}
		}

		value <- resp
	}()

	return
}

func (x xIO) ReadString(r io.Reader, delim byte) (value chan IOData[string]) {
	value = make(chan IOData[string], 1)

	go func() {
		var resp IOData[string]

		switch reader := r.(type) {
		case CYBReaderLike:
			resp.Data, resp.Error = reader.ReadString(delim)

		default:
			buff := bufio.NewReader(r)
			resp.Data, resp.Error = buff.ReadString('\n')

			if resp.Error == nil {
				resp.Data = strings.TrimSuffix(resp.Data, string([]byte{delim}))
			}
		}

		value <- resp
	}()

	return
}

func (x xIO) ReadStringLine(r io.Reader) (value chan IOData[string]) {
	value = make(chan IOData[string], 1)

	go func() {
		var resp IOData[string]

		switch reader := r.(type) {
		case CYBReaderLike:
			resp.Data, resp.Error = reader.ReadStringLine()

		default:
			scanner := bufio.NewScanner(r)
			scanner.Split(bufio.ScanLines)

			if scanner.Scan() {
				resp.Data = scanner.Text()
			}

			resp.Error = scanner.Err()
		}

		value <- resp
	}()

	return
}

var IO xIO
