package tests

import (
	"cyberpull.com/gokit/cyb"
)

func startCybServer(server *cyb.Server, opts cyb.Options) (err error) {
	opts.Info = cyb.Info{
		Name:        "Demo Server",
		Description: "CYB Demo Server",
	}

	server.Routes(addRoutes())

	err = <-server.Connect(&opts)

	if err != nil {
		return
	}

	go server.Run()

	return
}

func addRoutes() cyb.RequestRouterCallback {
	return func(router cyb.RequestRouter) {
		router.Set("GET", "/test/request", func(ctx *cyb.Context) cyb.Output {
			return ctx.Data("Demo Request Successful")
		})

		router.Set("GET", "/test/update", func(ctx *cyb.Context) cyb.Output {
			ctx.Update("Demo Update Successful")
			return ctx.Data("Demo Update Successful")
		})
	}
}
