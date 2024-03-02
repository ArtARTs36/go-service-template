package middlewares

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Log struct {
	handler http.Handler
}

func NewLog(handler http.Handler) *Log {
	return &Log{
		handler: handler,
	}
}

func (l *Log) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Debugf("handling request %s", req.RequestURI)

	l.handler.ServeHTTP(w, req)

	log.Debugf("request %s handled", req.RequestURI)
}
