package api

import (
	"net/http"
	"strconv"

	"github.com/archey347/goping/internal/api/routes"
	"github.com/archey347/goping/internal/config"
	"go.uber.org/zap"
)

func Start(c config.Config, l *zap.Logger) {
	r := routes.GetRouter()

	connectionBinding := ":" + strconv.Itoa(c.ApiServer.Port)

	l.Info("Starting http server with binding " + connectionBinding)

	http.ListenAndServe(connectionBinding, r)
}
