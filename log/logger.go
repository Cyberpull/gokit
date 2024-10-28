package log

import (
	"fmt"
	"log"
	"sync"

	"github.com/Cyberpull/gokit/color"
)

type Logger struct {
	*log.Logger

	mutex sync.Mutex
	color color.Attribute
}

func (x *Logger) SetColor(color color.Attribute) {
	x.mutex.Lock()
	defer x.mutex.Unlock()

	x.color = color
}

func (x *Logger) Fatal(a ...any) {
	x.Logger.Fatal(x.s(a...))
}

func (x *Logger) Fatalf(format string, a ...any) {
	x.Logger.Fatal(x.sf(format, a...))
}

func (x *Logger) Fatalln(a ...any) {
	x.Logger.Fatalln(x.s(a...))
}

func (x *Logger) Fatalfln(format string, a ...any) {
	x.Logger.Fatalln(x.sf(format, a...))
}

func (x *Logger) Panic(a ...any) {
	x.Logger.Panic(x.s(a...))
}

func (x *Logger) Panicf(format string, a ...any) {
	x.Logger.Panic(x.sf(format, a...))
}

func (x *Logger) Panicln(a ...any) {
	x.Logger.Panicln(x.s(a...))
}

func (x *Logger) Panicfln(format string, a ...any) {
	x.Logger.Panicln(x.sf(format, a...))
}

func (x *Logger) Print(a ...any) {
	x.Logger.Print(x.s(a...))
}

func (x *Logger) Printf(format string, a ...any) {
	x.Logger.Print(x.sf(format, a...))
}

func (x *Logger) Println(a ...any) {
	x.Logger.Println(x.s(a...))
}

func (x *Logger) Printfln(format string, a ...any) {
	x.Logger.Println(x.sf(format, a...))
}

// ====================================

func (x *Logger) s(v ...any) string {
	if x.color == 0 {
		return fmt.Sprint(v...)
	}

	x.mutex.Lock()
	defer x.mutex.Unlock()

	return color.String(x.color, v...)
}

func (x *Logger) sf(format string, v ...any) string {
	if x.color == 0 {
		return fmt.Sprintf(format, v...)
	}

	x.mutex.Lock()
	defer x.mutex.Unlock()

	return color.Stringf(x.color, format, v...)
}

// ====================================

func New(opts Options) *Logger {
	return &Logger{
		Logger: log.New(opts.out, opts.prefix, opts.flag),
		color:  opts.color,
	}
}

func Copy(x *Logger, attr ...color.Attribute) *Logger {
	result := &Logger{Logger: x.Logger}

	if len(attr) > 0 {
		result.SetColor(attr[0])
	}

	return result
}

func Color(attr ...color.Attribute) *Logger {
	result := Default()

	if len(attr) > 0 {
		result.SetColor(attr[0])
	}

	return result
}
