package api

import (
	"github.com/gofiber/fiber/v2"
)

type Version struct {
	Version string `json:"version"`
}

func GetVersion(c *fiber.Ctx) error {

	version := Version{Version: "123"}

	return c.JSON(&version)
}
