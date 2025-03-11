package config

import (
	"github.com/BurntSushi/toml"
)

// Config - струтктура для хранения конфигурации приложения
type Config struct {
	Server struct {
		Port string `toml:port` // port of server
	} `toml:"server"`
	Database struct {
		Host     string `toml:host` // database host etc
		Port     string `toml:host` // database host etc
		User     string `toml:host` // database host etc
		Password string `toml:host` // database host etc
		Name     string `toml:host` // database host etc
	}
}

// function to load config from the toml file
func LoadConfig(path string) (*Config, error) {
	var config Config // Создаем екземпляр структуры Конфиг
	// Декларируем Томл файл и сохраняем данные в структуру Конфиг
	_, err := toml.DecodeFile(path, &config)

	if err != nil {
		return nil, err // Если не удалось декодировать - возврат ошибки
	}
	return &config, nil // Возвращаем указатель на структуру Config / nil (нет ошибки)
}
