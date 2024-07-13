package graceful

import "context"

type initializable interface {
	initialize()
}

func d[T initializable](x T) T {
	x.initialize()
	return x
}

type Runner func(grace Grace)

type Grace interface {
	context.Context
}

func Run(callback ...Runner) {
	for _, fn := range callback {
		instance.Run(fn)
	}
}
