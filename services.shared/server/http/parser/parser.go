package parser

import "github.com/gofiber/fiber/v2"

func ParseBody() {}

func ParseQuery() {}

func ParsePrams() {}

func GetToken(c *fiber.Ctx) string {
	t := getTokenFromCookie(c)
	if t == "" {
		t = getTokenFromBearer(c)
	}
	return t

}

func getTokenFromCookie(c *fiber.Ctx) string {
	return c.Cookies("token")
}

func getTokenFromBearer(c *fiber.Ctx) string {
	b := c.Get("Authorization")
	if b == "" {
		return ""
	}
	return b[7:]

}
