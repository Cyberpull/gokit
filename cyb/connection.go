package cyb

type Connection interface {
	ReadBytes(delim byte) (b []byte, err error)
	ReadLine() (b []byte, err error)
	ReadString(delim byte) (s string, err error)
	ReadStringLine() (s string, err error)

	Write(p []byte) (n int, err error)
	WriteLine(p []byte) (n int, err error)
	WriteString(s string) (n int, err error)
	WriteStringLine(s string) (n int, err error)
}
