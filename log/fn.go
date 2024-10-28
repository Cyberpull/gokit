package log

import "log"

func Fatal(v ...any) {
	logger.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	logger.Fatalf(format, v...)
}

func Fatalln(v ...any) {
	logger.Fatalln(v...)
}

func Panic(a ...any) {
	logger.Panic(a...)
}

func Panicf(format string, a ...any) {
	logger.Panicf(format, a...)
}

func Panicln(a ...any) {
	logger.Panicln(a...)
}

func Print(a ...any) {
	logger.Print(a...)
}

func Printf(format string, a ...any) {
	logger.Printf(format, a...)
}

func Println(a ...any) {
	logger.Println(a...)
}

// =======================

func Default() *Logger {
	return logger
}

func Flags() int {
	return log.Flags()
}

// =======================

var logger *Logger

func init() {
	logger = &Logger{Logger: log.Default()}
}