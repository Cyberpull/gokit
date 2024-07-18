package gokit

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type IOData[T any] struct {
	Error error
	Data  T
}

type xIO struct {
	//
}

func (x xIO) Read(r io.Reader) (value chan IOData[[]byte]) {
	value = make(chan IOData[[]byte], 1)

	go func() {
		resp := IOData[[]byte]{Data: make([]byte, 0)}
		_, resp.Error = r.Read(resp.Data)
		value <- resp
	}()

	return
}

func (x xIO) ReadSingleByte(r io.Reader) (value chan IOData[byte]) {
	value = make(chan IOData[byte], 1)

	go func() {
		var resp IOData[byte]

		switch reader := r.(type) {
		case *bufio.Reader:
			resp.Data, resp.Error = reader.ReadByte()

		default:
			resp.Data, resp.Error = bufio.NewReader(r).ReadByte()
		}

		value <- resp
	}()

	return
}

func (x xIO) ReadSingleRune(r io.Reader) (value chan IOData[rune]) {
	value = make(chan IOData[rune], 1)

	go func() {
		var resp IOData[rune]

		switch reader := r.(type) {
		case *bufio.Reader:
			resp.Data, _, resp.Error = reader.ReadRune()

		default:
			resp.Data, _, resp.Error = bufio.NewReader(r).ReadRune()
		}

		value <- resp
	}()

	return
}

func (x xIO) ReadBytes(r io.Reader, delim byte) (value chan IOData[[]byte]) {
	value = make(chan IOData[[]byte], 1)

	go func() {
		var resp IOData[[]byte]

		switch reader := r.(type) {
		case *bufio.Reader:
			resp.Data, resp.Error = reader.ReadBytes(delim)

		default:
			buff := bufio.NewReader(r)
			resp.Data, resp.Error = buff.ReadBytes(delim)
		}

		resp.Data = bytes.TrimSuffix(resp.Data, []byte{delim})

		value <- resp
	}()

	return
}

func (x xIO) ReadLine(r io.Reader) (value chan IOData[[]byte]) {
	return x.ReadBytes(r, '\n')
}

func (x xIO) ReadString(r io.Reader, delim byte) (value chan IOData[string]) {
	value = make(chan IOData[string], 1)

	go func() {
		var resp IOData[string]

		switch reader := r.(type) {
		case *bufio.Reader:
			resp.Data, resp.Error = reader.ReadString(delim)

		default:
			resp.Data, resp.Error = bufio.NewReader(r).ReadString(delim)
		}

		resp.Data = strings.TrimSuffix(resp.Data, string([]byte{delim}))

		value <- resp
	}()

	return
}

func (x xIO) ReadLineString(r io.Reader) (value chan IOData[string]) {
	return x.ReadString(r, '\n')
}

var IO xIO
