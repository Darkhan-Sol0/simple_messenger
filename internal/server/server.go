package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simple_messenger/internal/config"
	"simple_messenger/internal/datasource/database"
	"simple_messenger/internal/service"
	"simple_messenger/internal/web"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	server struct {
		httpDriver *echo.Echo
		cfg        config.Config
		database   database.Database
		service    service.Service
		router     web.Routing
	}

	Server interface {
		Run() (err error)
	}
)

func New() Server {
	return &server{
		httpDriver: echo.New(),
		cfg:        config.GetConfig(),
	}
}

func (s *server) start() {
	log.Printf("Starting server at %s\n", s.cfg.GetAddress())
	if err := s.httpDriver.Start(s.cfg.GetAddress()); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func (s *server) shutdown(ctx context.Context) {
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	log.Printf("Shutting down the server...")
	if err := s.httpDriver.Shutdown(shutdownCtx); err != nil {
		log.Printf("Graceful shutdown failed with error: %v\n", err)
	}
	log.Println("Server gracefully stopped.")
}

func (s *server) Run() error {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	err := s.connect(ctx)
	if err != nil {
		return err
	}
	defer s.closeDatabase(ctx)
	go func() {
		s.start()
		stop()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("Context canceled, stopping server...")
	case <-quit:
		log.Println("Received termination signal, stopping server...")
	}
	s.shutdown(ctx)
	return nil
}

func (s *server) connect(ctx context.Context) error {
	err := s.initDatabase(ctx)
	if err != nil {
		return err
	}
	s.connectService()
	s.connectWeb()
	return nil
}

func (s *server) initDatabase(ctx context.Context) (err error) {
	s.database, err = database.New(ctx, s.cfg)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) closeDatabase(ctx context.Context) {
	s.database.Close(ctx)
}

func (s *server) connectService() {
	s.service = service.New(s.database)
}

func (s *server) connectWeb() {
	s.router = web.New(s.service)
	s.router.RegisterRoutes(s.httpDriver)
}
