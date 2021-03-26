package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"io/ioutil"
	"fmt"
	"github.com/google/uuid"
	// "reflect"
)

type Name struct {
	Name string
}

const (
	host			= "localhost"
	port 			= 5432
	user 			= "postgres"
	password 	= "1234"
	dbname 		= "super"
)

func postApi(c *fiber.Ctx) error {
	type request struct {
		Name string
	}
	
	var body request

	err := c.BodyParser(&body)
	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newName := Name{
		Name: body.Name,
	}
	
	res, err := http.Get("https://www.superheroapi.com/api.php/4329276143753700/search/" + newName.Name)

	if err != nil {
		log.Fatal( err )
	}

	data, _ := ioutil.ReadAll( res.Body )
	
	res.Body.Close()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	INSERT INTO public.superhero (name, fullname, intelligence, power, occupation, image, uuid)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`

	uuid := uuid.New().String()

	id:= 0
	err = db.QueryRow(sqlStatement, "john", "john smith", "38", "34", "teacher", "..", uuid).Scan(&id)
	
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}

	fmt.Println("New record ID is", id)

	fmt.Printf("%s\n", data)
	
	// fmt.Printf(reflect.TypeOf(data).String())


	return c.SendString("Hello, World")
}

func main() {
    app := fiber.New()
		
		app.Use(logger.New())

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("ola")
		})

    app.Post("/new", postApi)

    app.Listen(":3000")
}
