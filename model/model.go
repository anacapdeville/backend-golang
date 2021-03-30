package model

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/google/uuid"
	"reflect"
	"github.com/gofiber/fiber/v2"
)

type Hero struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
	Fullname string `json:"Fullname"`
	Intelligence string `json:"Intelligence"`
	Power string `json:"Power"`
	Occupation string `json:"Occupation"`
	Image string `json:"Image"`
	Group_affiliation interface{} `json:"Group_affiliation"`
	Number_relatives interface{} `json:"Number_relatives"`
	Uuid interface{} `json:"Uuid"`
}

type Request struct {
	Name string
}

func GetAll () ([]Hero, error) {
	rows, err := Connection().Query("SELECT * FROM superhero")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var heros []Hero

	for rows.Next() {
		var id int
		var name string
		var fullname string
		var intelligence string
		var power string
		var occupation string
		var image string
		var group_affiliation interface{}
		var number_relatives interface{}
		var uuid interface{}

		err = rows.Scan(&id, &name, &fullname, &intelligence, &power, &occupation, &image, &group_affiliation, &number_relatives, &uuid)
		
		if err != nil {
			return nil, err
		}

		hero := Hero{id, name, fullname, intelligence, power, occupation, image, group_affiliation, number_relatives, uuid}

		heros = append(heros, hero)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return heros, nil
}

func GetByName (name string) (*Hero, error) {
	sqlStatement := `SELECT * FROM superhero WHERE name=$1`

	super := &Hero{}

	row := Connection().QueryRow(sqlStatement, name)

	err := row.Scan(&super.Name, &super.Fullname, &super.Intelligence, &super.Power, &super.Occupation, &super.Image, &super.Group_affiliation, &super.Number_relatives, &super.Uuid)

	fmt.Println(super)

	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("Super não encontrado")
	case nil:
		return super, nil
	default:
		return nil, err
	}
}

func AddSuper (c *fiber.Ctx) error {	
	var body Request

	err := c.BodyParser(&body)
	
	if err != nil {
		return err
	}

	newName := Request{
		Name: body.Name,
	}
	
	res, err := http.Get("https://www.superheroapi.com/api.php/4329276143753700/search/" + newName.Name)

	if err != nil {
		return err
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
	INSERT INTO superhero (name, fullname, intelligence, power, occupation, image, uuid)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING uuid`

	uuid := uuid.New().String()

	err = db.QueryRow(sqlStatement, "carla", "carla smith", "38", "34", "teacher", "..", uuid).Scan(&uuid)
	
	if err != nil {
		return err
	}
	
	fmt.Printf("%s\n", data)
	
	fmt.Printf(reflect.TypeOf(data).String())

	return c.Status(fiber.StatusOK).JSON("Uuid inserted:" + uuid)
}
