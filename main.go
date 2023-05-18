package main

import (
	"net/http"

	"go-rest-i18n/pkg/locale"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
	locale.New()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		lang := c.GetReqHeaders()["Accept-Language"]
		localizer := i18n.NewLocalizer(locale.Bundle, lang)

		resp := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "HelloWorld",
		})
		return c.Status(http.StatusOK).JSON(map[string]any{
			"data": resp,
		})

	})
	app.Listen(":3000")
}
