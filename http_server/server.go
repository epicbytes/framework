package http_server

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App    *fiber.App
	Config *Config
	Done   chan struct{}
}

func NewServer(config *Config, handler fiber.ErrorHandler) *Server {

	cfg := fiber.Config{
		ServerHeader: "EpicServer",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	}
	if handler != nil {
		cfg.ErrorHandler = handler
	}
	app := fiber.New(cfg)

	server := &Server{
		App:    app,
		Config: config,
		Done:   make(chan struct{}),
	}

	return server
}

func (s *Server) StartServer() error {
	return s.App.Listen(s.Config.Address)
}

func (s *Server) StopServer() error {
	server := s.App
	return server.Shutdown()
}
