package error_handler

import "github.com/gofiber/fiber/v2"

func New() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		// TODO: Add your own error handling
		return err
	}
}
