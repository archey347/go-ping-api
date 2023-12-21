package api

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/archey347/goping/internal/api/routes"
	"github.com/archey347/goping/internal/config"
	"github.com/coreos/go-systemd/daemon"
	"go.uber.org/zap"
)

func Start(c config.Config, l *zap.Logger) {
	router := routes.GetRouter()

	connectionBinding := c.ApiServer.Address + ":" + strconv.Itoa(c.ApiServer.Port)

	l.Info("Starting http server with binding " + connectionBinding)

	listener, err := net.Listen("tcp", connectionBinding)

	if err != nil {
		l.Fatal("Can't listen: " + err.Error())
		return
	}

	daemon.SdNotify(false, daemon.SdNotifyReady)
	go watchdog(c)

	http.Serve(listener, router)
}

func watchdog(c config.Config) {
	var url string

	if c.ApiServer.Address == "" {
		url = "http://127.0.0.1:" + strconv.Itoa(c.ApiServer.Port)
	} else {
		url = "http://" + c.ApiServer.Address + ":" + strconv.Itoa(c.ApiServer.Port) + "/ping"
	}

	for {
		_, err := http.Get(url)

		if err == nil {
			daemon.SdNotify(false, daemon.SdNotifyWatchdog)
		}

		time.Sleep(5 * time.Second)
	}
}
