package tests

import (
	"time"

	"cyberpull.com/gokit/cyb"
)

func startCybClient(client *cyb.Client, address string) (err error) {
	err = <-client.Connect(&cyb.Options{
		Network:    "unix",
		SocketPath: address,
		Info: cyb.Info{
			Name:        "Demo Client",
			Description: "CYB Demo Client",
		},
	})

	if err != nil {
		return
	}

	go client.Run()

	time.Sleep(time.Second / 2)

	return
}
