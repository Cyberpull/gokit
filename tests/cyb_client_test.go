package tests

import "cyberpull.com/gokit/cyb"

func startCybClient(client *cyb.Client, address string) (err error) {
	client.Options(&cyb.Options{
		Network:    "unix",
		SocketPath: address,
		Info: cyb.Info{
			Name:        "Demo Client",
			Description: "CYB Demo Client",
		},
	})

	return <-client.Start(false)
}
