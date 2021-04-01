package main

import (
	_ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/anacapdeville/backend-golang/model"
)

type Name struct {
	Name string
}

type Uuid struct {
	Uuid string
}

func main() {
  app := fiber.New()
		
	app.Use(logger.New())

	app.Post("/new", model.AddSuper)

	app.Get("/", func (c *fiber.Ctx) error {
		allSupers, err := model.GetAll()
		
		if err != nil {
			return err
		}
		
		return c.Status(fiber.StatusOK).JSON(allSupers)
	})

	app.Get("/getbyname/:name", func (c *fiber.Ctx) error {

		name := c.Params("name")
		
		super, err := model.GetByName(name)
		
		if err != nil {
			return err
		}
			
		return c.Status(fiber.StatusOK).JSON(super)
	})

	app.Get("/getbyuuid/:uuid", func (c *fiber.Ctx) error {
		uuid := c.Params("uuid")
		
		super, err := model.GetByUuid(uuid)
		
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(super)
	})

	app.Delete("/delete/:uuid", model.DeleteSuper)

  app.Listen(":3000")
}
