package errors

type Error struct {
	code    int
	message string
	stack   string
}

func (x Error) Code() int {
	return x.code
}

func (x Error) Error() string {
	return x.String()
}

func (x Error) Stack() string {
	return x.stack
}

func (x Error) String() string {
	message := x.message

	if x.stack != "" {
		message += "\n" + x.stack
	}

	return message
}
