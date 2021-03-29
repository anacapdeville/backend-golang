package main

import (
	// "database/sql"
	_ "github.com/lib/pq"
	// "net/http"
	// "log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "io/ioutil"
	// "fmt"
	// "github.com/google/uuid"
	// "reflect"
	"github.com/anacapdeville/backend-golang/model"
)

type Name struct {
	Name string
}

// const (
// 	host			= "localhost"
// 	port 			= 5432
// 	user 			= "postgres"
// 	password 	= "1234"
// 	dbname 		= "super"
// )

// func postApi(c *fiber.Ctx) error {
// 	type request struct {
// 		Name string
// 	}
	
// 	var body request

// 	err := c.BodyParser(&body)
	
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Cannot parse JSON",
// 		})
// 	}

// 	newName := Name{
// 		Name: body.Name,
// 	}
	
// 	res, err := http.Get("https://www.superheroapi.com/api.php/4329276143753700/search/" + newName.Name)

// 	if err != nil {
// 		log.Fatal( err )
// 	}

// 	data, _ := ioutil.ReadAll( res.Body )
	
// 	res.Body.Close()

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	sqlStatement := `
// 	INSERT INTO public.superhero (name, fullname, intelligence, power, occupation, image, uuid)
// 	VALUES ($1, $2, $3, $4, $5, $6, $7)
// 	RETURNING id`

// 	uuid := uuid.New().String()

// 	id:= 0
// 	err = db.QueryRow(sqlStatement, "john", "john smith", "38", "34", "teacher", "..", uuid).Scan(&id)
	
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("New record ID is", id)

// 	fmt.Printf("%s\n", data)
	
// 	fmt.Printf(reflect.TypeOf(data).String())


// 	return c.SendString("Hello, World")
// }

// func getAll(c *fiber.Ctx) error {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)

// 	rows, err := db.Query("SELECT * FROM superhero")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		var fullname string
// 		var intelligence string
// 		var power string
// 		var occupation string
// 		var image string
// 		var group_affiliation interface{}
// 		var number_relatives interface{}
// 		var uuid interface{}

// 		err = rows.Scan(&id, &name, &fullname, &intelligence, &power, &occupation, &image, &group_affiliation, &number_relatives, &uuid)
		
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(id, name, fullname, intelligence, power, occupation, image,group_affiliation, number_relatives, uuid)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c.SendString("Hello, World")
// }

// func getByName(c *fiber.Ctx) error {
// 	type request struct {
// 		Name string
// 	}
	
// 	var body request

// 	err := c.BodyParser(&body)

// 	nameOf := Name{
// 		Name: body.Name,
// 	}
	
// 	type Super struct {
// 		id int
// 		name string
// 		fullname string
// 		intelligence string
// 		power string
// 		occupation string
// 		image string
// 		group_affiliation interface{}
// 		number_relatives interface{}
// 		uuid interface{}
// 	}

// 	sqlStatement := `SELECT * FROM superhero WHERE name=$1`

// 	var super Super

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)

// 	row := db.QueryRow(sqlStatement, nameOf.Name)

// 	err = row.Scan(&super.id, &super.name, &super.fullname, &super.intelligence, &super.power, &super.occupation, &super.image, &super.group_affiliation, &super.number_relatives, &super.uuid)

// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 		return c.SendString("No rows")
// 	case nil:
// 		fmt.Println(super)
// 	default:
// 		panic(err)
// 	}
// 	return c.SendString("Hello, World")
// }

func getByName(c *fiber.Ctx) error {
	var body Name

	err := c.BodyParser(&body)

	if err != nil {
		panic(err)
	}
	
	nameOf := Name{
		Name: body.Name,
	}

	super, err := model.GetByName(nameOf.Name)

	if err != nil {
		panic(err)
	}
	return super
}

func main() {
    app := fiber.New()
		
		app.Use(logger.New())

		app.Get("/", model.GetAll)

		app.Get("/getbyname", getByName)

    // app.Post("/new", postApi)

    app.Listen(":3000")
}
