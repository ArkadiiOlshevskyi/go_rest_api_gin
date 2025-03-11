package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// init Handlers - функция для инициализации обработчиков маршрутов
func InitHandlers(router *gin.Engine) {
	router.GET("/ping", PingHandler)      // Регистрация обработчка GET запросов
	router.POST("/data", PostDataHandler) // Регистрация обработчка POST запросов
}

// pingHandler - обработчик GET запроса на /ping
func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) // Возвращаем JSON-ответ "pong"
}

// pongHandler - обработчик POST запроса на /dataб т.е. то что вернул нам сервер
func PostDataHandler(ctx *gin.Context) {
	var data map[string]interface{} // обявляем переменную для хранения данных

	// Пытаемся парсить JSON -запрос и сохранить данные в переменную data
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Если произошла ошибка, возвращаем HTTP-статус 400 и сообщение об ошибке
		return
	}

	// Возвращаем полученные данные в JSON-ответе
	ctx.JSON(http.StatusOK, gin.H{"received": data}) // Возвращаем JSON-ответ с полученными данными
}
