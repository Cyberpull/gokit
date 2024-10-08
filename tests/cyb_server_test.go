package tests

import (
	"github.com/Cyberpull/gokit/cyb"
)

func startCybServer(server *cyb.Server, socket string) (err error) {
	opts := &cyb.Options{
		Socket: socket,
		Info: cyb.Info{
			Name:        "Demo Server",
			Description: "CYB Demo Server",
		},
	}

	server.Routes(addRoutes())

	err = <-server.Connect(opts)

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

		router.Set("GET", "/test/error", func(ctx *cyb.Context) cyb.Output {
			return ctx.Error("Demo Error Successful")
		})

		router.Set("GET", "/test/struct", func(ctx *cyb.Context) cyb.Output {
			return ctx.Data(DemoResponse{
				Name:  "Christian",
				Email: "demo@example.com",
			})
		})

		router.Set("GET", "/test/struct/update", func(ctx *cyb.Context) cyb.Output {
			ctx.Update(DemoResponse{
				Name:  "Christian",
				Email: "demo@example.com",
			})

			return ctx.Data("Struct Update Successful")
		})

		router.Set("GET", "/test/struct/request", func(ctx *cyb.Context) cyb.Output {
			var data DemoRequest

			if err := ctx.Bind(&data); err != nil {
				return ctx.Error(err)
			}

			return ctx.Data("Success!")
		})
	}
}
