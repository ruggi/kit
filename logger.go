package kit

import (
	"net/http"
	"os"
	"time"

	logging "github.com/op/go-logging"
)

type Logger struct {
	*logging.Logger
}

var defaultBackend = logging.AddModuleLevel(
	logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stdout, "", 0),
		logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{module} %{level:.4s} ▶ %{id:03x}%{color:reset} %{message}"),
	))

func NewLogger(module string) *Logger {
	l := &Logger{logging.MustGetLogger(module)}
	l.SetBackend(defaultBackend)
	return l
}

func (l *Logger) LogRequest(r *http.Request) {
	l.Infof("%s %s", r.Method, r.URL)
}

func (l *Logger) LogSuccess(r *http.Request, dur time.Duration, status int, resp interface{}) {
	l.Noticef("%s %s (%s) (%d)", r.Method, r.URL, dur, status)
}

func (l *Logger) LogError(r *http.Request, status int, err error) {
	l.Errorf("%s %s: %s (%d)", r.Method, r.URL, err, status)
}
