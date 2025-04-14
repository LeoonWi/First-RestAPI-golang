package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	messages     = make(map[int]string)
	index    int = 0
	mu       sync.Mutex
)

func getMessage(c echo.Context) error {
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

	index++
	messages[index] = message.Text

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

	// блокировка Mutex необходима для корректной работы с картами
	// горутины могут начать date race
	mu.Lock()
	_, k := messages[id] // value, key := message[id] - value может быть "", key будет True или False
	mu.Unlock()
	if !k {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Unknown ID",
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

	messages[id] = message.Text
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

	mu.Lock()
	_, k := messages[id]
	mu.Unlock()
	if !k {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Unknown ID",
		})
	}

	delete(messages, id)
	return c.JSON(http.StatusOK, Response{
		Status:  "Done",
		Message: "Message has been deleted",
	})
}

func main() {
	e := echo.New()
	e.GET("/", getMessage)
	e.POST("/", postMessage)
	e.PATCH("/:id", patchMessage)
	e.DELETE("/:id", deleteMessage)
	e.Start(":8080")
}
