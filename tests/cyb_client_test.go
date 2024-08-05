package tests

import (
	"time"

	"github.com/Cyberpull/gokit/cyb"
)

func startCybClient(client *cyb.Client, socket string) (err error) {
	opts := &cyb.Options{
		Socket: socket,
		Info: cyb.Info{
			Name:        "Demo Client",
			Description: "CYB Demo Client",
		},
	}

	client.SetRequestTimeout(10)

	err = <-client.Connect(opts)

	if err != nil {
		return
	}

	go client.Run()

	time.Sleep(time.Second / 2)

	return
}
