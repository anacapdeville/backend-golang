package main

import (
	// "database/sql"
	_ "github.com/lib/pq"
	// "net/http"
	// "log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "io/ioutil"
	"fmt"
	// "github.com/google/uuid"
	// "reflect"
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

		app.Get("/", func (c *fiber.Ctx) error {
			allSupers, err := model.GetAll()
		
			fmt.Println(allSupers)
		
			if err != nil {
				return err
			}
		
			return c.Status(fiber.StatusOK).JSON(allSupers)
		})

		app.Get("/getbyname/:name", func (c *fiber.Ctx) error {
			// var body Name
		
			// err := c.BodyParser(&body)
		
			// if err != nil {
			// 	return err
			// }
			
			// nameOf := Name{
			// 	Name: body.Name,
			// }

			name := c.Params("name")
		
			super, err := model.GetByName(name)
		
			if err != nil {
				return err
			}
			return c.Status(fiber.StatusOK).JSON(super)
		})

		app.Get("/getbyuuid/:uuid", func (c *fiber.Ctx) error {
			// var body Uuid
		
			// err := c.BodyParser(&body)
		
			// if err != nil {
			// 	return err
			// }
			
			// uuidOf := Uuid{
			// 	Uuid: body.Uuid,
			// }

			uuid := c.Params("uuid")
		
			super, err := model.GetByName(uuid)
		
			if err != nil {
				return err
			}
			return c.Status(fiber.StatusOK).JSON(super)
		})

    app.Post("/new", model.AddSuper)

    app.Listen(":3000")
}
