package parser

import (
	"clean-boilerplate/shared/i18n"
	i18nHttp "clean-boilerplate/shared/server/http/i18n"
	"clean-boilerplate/shared/server/http/result"
	"clean-boilerplate/shared/validator"

	"github.com/gofiber/fiber/v2"
)

func ParseBody(c *fiber.Ctx, v validator.Validator, i i18n.I18n, d interface{}) {
	l, a := i18nHttp.GetLanguagesInContext(i, c)
	if err := c.BodyParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_body", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func ParseQuery(c *fiber.Ctx, v validator.Validator, i i18n.I18n, d interface{}) {
	l, a := i18nHttp.GetLanguagesInContext(i, c)
	if err := c.QueryParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_query", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func ParseParams(c *fiber.Ctx, v validator.Validator, i i18n.I18n, d interface{}) {
	l, a := i18nHttp.GetLanguagesInContext(i, c)
	if err := c.ParamsParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_params", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func validateStruct(d interface{}, v validator.Validator, i i18n.I18n, l, a string) {
	if errors := v.ValidateStruct(d, l, a); len(errors) > 0 {
		panic(result.ErrorDetail(i.Translate("error_validation_failed", l, a), errors, fiber.StatusBadRequest))
	}
}

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
