package main

import (
	"github.com/first-restapi-golang/internal/handler"
	"github.com/first-restapi-golang/internal/model"
	"github.com/first-restapi-golang/internal/repository"
	"github.com/first-restapi-golang/internal/service"
	"github.com/labstack/echo/v4"
)

// func initDB() {
// 	dsn := "host=localhost user=postgres password=admin dbname=mydb port=5432 sslmode=disable"
// 	var err error
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
// 	}

// 	db.AutoMigrate(&Message{})
// }

// func getMessage(c echo.Context) error {
// 	var messages []Message

// 	if err := db.Find(&messages).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: fmt.Sprintf("Failed to receive messages: %v", err),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, &messages)
// }

// func postMessage(c echo.Context) error {
// 	var message Message
// 	if err := c.Bind(&message); err != nil {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: "Could not add the message",
// 		})
// 	}

// 	if message.Text == "" {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: "Text message could not empty",
// 		})
// 	}

// 	if err := db.Create(&message).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: "Could not create the message",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, Response{
// 		Status:  "Done",
// 		Message: "Message has been added",
// 	})
// }

// func patchMessage(c echo.Context) error {
// 	// Param отличается от QueryParam наличием ключей в URL
// 	// Param - :8080/3
// 	// QueryParam - :8080/?id=3
// 	idParam := c.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: "Bad ID",
// 		})
// 	}

// 	var message Message
// 	if err := c.Bind(&message); err != nil {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: "Could not add the message",
// 		})
// 	}

// 	if message.Text == "" {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: "Text message could not empty",
// 		})
// 	}

// 	if err := db.Model(&message).Where("id = ?", id).Update("text", message.Text).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: fmt.Sprintf("Could not update the message: %v", err),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, Response{
// 		Status:  "Done",
// 		Message: "Message has been updated",
// 	})

// }

// func deleteMessage(c echo.Context) error {
// 	idParam := c.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: "Bad ID",
// 		})
// 	}

// 	if err := db.Delete(&Message{}, id).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, Response{
// 			Status:  "Error",
// 			Message: fmt.Sprintf("Could not delete the message: %v", err),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, Response{
// 		Status:  "Done",
// 		Message: "Message has been deleted",
// 	})
// }

func main() {
	db := repository.InitDB("host=db user=postgres password=admin dbname=mydb port=5432 sslmode=disable")
	db.AutoMigrate(&model.Category{})
	e := echo.New()

	r := repository.New(db)
	s := service.New(r)
	handler.New(e, s)

	e.Start(":8080")
}
