package gameserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GameServer ...
type GameServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *GameServer {
	return &GameServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *GameServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Starting FatalForce game server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

//configureLogger ...
func (s *GameServer) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

//configureRouter ...
func (s *GameServer) configureRouter() {
	s.router.HandleFunc("/test", s.handleTest())
}

func (s *GameServer) handleTest() http.HandlerFunc {
	///// ....

	return func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Test Request!")
	}
}
