package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorestapi/config"
	"gorestapi/internal/rest"
)

func main() {
	// 1 - Загрузка конфига из файла config.toml
	cfg, err := config.LoadConfig("./config/config.toml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err) // Если не удалось загрузить конфигурацию, завершаем программу
	}

	// 2 - Инициализация сервера с загруженной конфигом
	server := rest.NewServer(cfg)

	// 3 - Запуск Сервера в отдельной горутине, чтобы не блокировать основной канал
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err) // Если сервер не запустился, завершаем программу
		}
	}()

	log.Println("Server started")

	// 4 - gracefull shutdown: обработка сигналов SIGINT SIGTERM для корректной остановки
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // Подписываемся на сигналы
	<-quit                                               // Ожидаем получения сигнала
	log.Println("Server shutting down...")               // Выводим сообщение о начале завершения работы

	// 5 - Создание контекста с таймаутом для завершения работа сервера
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Гарантируем отмену контекста после завершения функции

	// 6 - Остановка сервера с искользованием контекста
	if err := server.Stop(ctx); err != nil {
		log.Fatal("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
