package rotes

import (
	"go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	// /api/v1
	// api
	api := app.Group("/api")
	//v1
	v1 := api.Group("/v1")
	v1.Get("/", controllers.HelloTest)
	v1.Post("/", controllers.BodyParserTest)
	v1.Get("/user/:name", controllers.ParamsTest)
	v1.Post("/inet", controllers.QueryTest)
	v1.Post("/valid", controllers.ValidateTest)

	v2 := api.Group("/v2")
	v2.Get("/", controllers.HelloTestV2)

	//client
	//400
	//401 unAuthorized
	//403 access denin
	//404 record not found

	//
	//200 success
	//201 create success

	//server
	//500 internal server error
	//502
	//503
}
