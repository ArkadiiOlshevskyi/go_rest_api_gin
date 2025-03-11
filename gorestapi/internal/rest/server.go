package rest

import (
	"context"
	"net/http"

	"gorestapi/config"

	"github.com/gin-gonic/gin"
)

// Server - структура для хранения информации о сервере
type Server struct {
	httpServer *http.Server
	config     *config.Config
}

// NewServer - конструктор для создания нового серверва екземпляр Сервер
func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

// Start - метод запуска сервера
func (s *Server) Start() error {
	gin.SetMode(gin.ReleaseMode) // Устанавливаем режим релиза
	router := gin.Default()      // Создаем новый екземпляр роутера на GIN

	// Инициализация обработчиковб передаем роутер GIN
	handlers.InitHandlers(router)

	// Настройка HTTP-сервера
	s.httpServer = &http.Server{
		Addr:    ":" + s.config.Server.Port, // Адрес сервера из конфигурации
		Handler: router,                     // Роутер GIN в качестве обработчика
	}

	//
	return s.httpServer.ListenAndServe()
}

// Stop - метод для остановки сервера
func (s *Server) Stop(ctx context.Context) error {
	// Остановка HTTP-сервера с использованием контекста для таймаута
	return s.httpServer.Shutdown(ctx)
}

// Status - метод для проверки статуса сервера
func (s *Server) Status() string {
	// Реализация проверки статуса сервера (например, проверка подключения к базе данных)
	return "OK"
}
