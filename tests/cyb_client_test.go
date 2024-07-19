package tests

import (
	"time"

	"cyberpull.com/gokit/cyb"
)

func startCybClient(client *cyb.Client, opts cyb.Options) (err error) {
	opts.Info = cyb.Info{
		Name:        "Demo Client",
		Description: "CYB Demo Client",
	}

	err = <-client.Connect(&opts)

	if err != nil {
		return
	}

	go client.Run()

	time.Sleep(time.Second / 2)

	return
}
