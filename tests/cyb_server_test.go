package tests

import "cyberpull.com/gokit/cyb"

func startCybServer(server *cyb.Server, address string) (err error) {
	server.Options(&cyb.Options{
		Network:    "unix",
		SocketPath: address,
		Info: cyb.Info{
			Name:        "Demo Server",
			Description: "CYB Demo Server",
		},
	})

	server.Routes(addRoutes())

	return <-server.Listen()
}

func addRoutes() cyb.RequestRouterCallback {
	return func(router cyb.RequestRouter) {
		router.Set("GET", "/test/request", func(ctx *cyb.Context) cyb.Output {
			return ctx.Data("Demo Request Successful")
		})

		router.Set("GET", "/test/update", func(ctx *cyb.Context) cyb.Output {
			ctx.Update(cyb.Data{
				Code:    200,
				Content: "Demo Update Successful",
			})

			return ctx.Data("Demo Update Successful")
		})
	}
}
