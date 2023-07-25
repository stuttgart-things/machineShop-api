package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type VM struct {
	Name string `json:"name"`
	Os   string `json:"os"`
}

func CreateVM(c *fiber.Ctx) error {

	vm := VM{}

	if err := c.BodyParser(&vm); err != nil {
		return err
	}

	fmt.Println(vm)

	return c.JSON(vm)
}
