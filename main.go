package main

import (
	"flag"

	"github.com/golang/glog"

	"github.com/gwangyi/RachelHomeService/api"
)

func main() {
	flag.Parse()

	service, err := api.CreateService()

	if err != nil {
		glog.Fatal(err)
		return
	}

	glog.Fatal(service.Start())
}
