package dbo

type initializable interface {
	initialize()
}

func d[T initializable](x T) T {
	x.initialize()
	return x
}
