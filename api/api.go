package api

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

type RachelHomeService struct {
	echo *echo.Echo
}

var port = flag.Int("port", 5409, "HTTP Listening Port")

func CreateService() (*RachelHomeService, error) {
	service := &RachelHomeService{}

	service.echo = echo.New()
	service.echo.GET("/quitquitquit", func(c echo.Context) error {
		glog.Infof("Received quitquitquit")
		service.echo.Shutdown(context.Background())
		return nil
	})
	return service, nil
}

func (service *RachelHomeService) Start() error {
	glog.V(2).Infof("Starting server at :%d", *port)
	return service.echo.Start(fmt.Sprintf(":%d", *port))
}
