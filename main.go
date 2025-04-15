package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Text string `gorm:"type:text" json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=admin dbname=mydb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	db.AutoMigrate(&Message{})
}

func getMessage(c echo.Context) error {
	var messages []Message

	if err := db.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: fmt.Sprintf("Failed to receive messages: %v", err),
		})
	}

	return c.JSON(http.StatusOK, &messages)
}

func postMessage(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	if message.Text == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Text message could not empty",
		})
	}

	if err := db.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not create the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Done",
		Message: "Message has been added",
	})
}

func patchMessage(c echo.Context) error {
	// Param отличается от QueryParam наличием ключей в URL
	// Param - :8080/3
	// QueryParam - :8080/?id=3
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	if message.Text == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Text message could not empty",
		})
	}

	if err := db.Model(&message).Where("id = ?", id).Update("text", message.Text).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: fmt.Sprintf("Could not update the message: %v", err),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Done",
		Message: "Message has been updated",
	})

}

func deleteMessage(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	if err := db.Delete(&Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: fmt.Sprintf("Could not delete the message: %v", err),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Done",
		Message: "Message has been deleted",
	})
}

func main() {
	initDB()

	e := echo.New()
	e.GET("/", getMessage)
	e.POST("/", postMessage)
	e.PATCH("/:id", patchMessage)
	e.DELETE("/:id", deleteMessage)

	e.Start(":8080")
}
