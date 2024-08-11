package handler

import "github.com/gofiber/fiber/v3"

type WebHandler struct {
}

func NewWebhandler() *WebHandler {
	return &WebHandler{}
}

func (wh *WebHandler) Index(c *fiber.Ctx) {

}
