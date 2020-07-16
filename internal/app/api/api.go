package api

import (
	"github.com/gorilla/mux"
	"github.com/iluxaorlov/rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Api struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *Api {
	return &Api{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),

	}
}

func (a *Api) Start() error {
	if err := a.configLogger(); err != nil {
		return err
	}

	a.configRouter()

	if err := a.configStore(); err != nil {
		return err
	}

	a.logger.Info("Starting API server")

	return http.ListenAndServe(a.config.Address, a.router)
}

func (a *Api) configLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)

	return nil
}

func (a *Api) configRouter() {
	a.router.HandleFunc("/hello", a.handleHello())
}

func (a *Api) configStore() error {
	s := store.New(a.config.Store)

	if err := s.Open(); err != nil {
		return err
	}

	a.store = s

	return nil
}

func (a *Api) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	}
}