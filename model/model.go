package model

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/google/uuid"
	// "reflect"
	"github.com/gofiber/fiber/v2"
	"encoding/json"
	"strings"
)

type Hero struct {
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
		var name string
		var fullname string
		var intelligence string
		var power string
		var occupation string
		var image string
		var group_affiliation interface{}
		var number_relatives interface{}
		var uuid interface{}

		err = rows.Scan(&uuid, &name, &fullname, &intelligence, &power, &occupation, &image, &group_affiliation, &number_relatives)
		
		if err != nil {
			return nil, err
		}

		hero := Hero{name, fullname, intelligence, power, occupation, image, group_affiliation, number_relatives, uuid}

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

	err := row.Scan(&super.Uuid, &super.Name, &super.Fullname, &super.Intelligence, &super.Power, &super.Occupation, &super.Image, &super.Group_affiliation, &super.Number_relatives)

	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("Super não encontrado")
	case nil:
		return super, nil
	default:
		return nil, err
	}
}

func GetByUuid (uuid string) (*Hero, error) {
	sqlStatement := `SELECT * FROM superhero WHERE uuid=$1`

	super := &Hero{}

	row := Connection().QueryRow(sqlStatement, uuid)

	err := row.Scan(&super.Uuid, &super.Name, &super.Fullname, &super.Intelligence, &super.Power, &super.Occupation, &super.Image, &super.Group_affiliation, &super.Number_relatives)
	
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

	heroFound, err := GetByName(newName.Name)

	if heroFound != nil {
		return errors.New("Super já cadastrado")
	}
	
	res, err := http.Get("https://www.superheroapi.com/api.php/4329276143753700/search/" + newName.Name)

	if err != nil {
		return err
	}

	data, _ := ioutil.ReadAll( res.Body )
	
	res.Body.Close()

	type Result struct {
		Response   string `json:"response"`
		ResultsFor string `json:"results-for"`
		Results    []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Powerstats struct {
				Intelligence string `json:"intelligence"`
				Strength     string `json:"strength"`
				Speed        string `json:"speed"`
				Durability   string `json:"durability"`
				Power        string `json:"power"`
				Combat       string `json:"combat"`
			} `json:"powerstats"`
			Biography struct {
				FullName        string   `json:"full-name"`
				AlterEgos       string   `json:"alter-egos"`
				Aliases         []string `json:"aliases"`
				PlaceOfBirth    string   `json:"place-of-birth"`
				FirstAppearance string   `json:"first-appearance"`
				Publisher       string   `json:"publisher"`
				Alignment       string   `json:"alignment"`
			} `json:"biography"`
			Appearance struct {
				Gender    string   `json:"gender"`
				Race      string   `json:"race"`
				Height    []string `json:"height"`
				Weight    []string `json:"weight"`
				EyeColor  string   `json:"eye-color"`
				HairColor string   `json:"hair-color"`
			} `json:"appearance"`
			Work struct {
				Occupation string `json:"occupation"`
				Base       string `json:"base"`
			} `json:"work"`
			Connections struct {
				GroupAffiliation string `json:"group-affiliation"`
				Relatives        string `json:"relatives"`
			} `json:"connections"`
			Image struct {
				URL string `json:"url"`
			} `json:"image"`
		} `json:"results"`
	}

	var dataStruct Result

	if err := json.Unmarshal(data, &dataStruct); err != nil {
		return err
	}

	name := dataStruct.Results[0].Name
	fullname := dataStruct.Results[0].Biography.FullName
	intelligence := dataStruct.Results[0].Powerstats.Intelligence
	power := dataStruct.Results[0].Powerstats.Power
	occupation := dataStruct.Results[0].Work.Occupation
	image := dataStruct.Results[0].Image.URL
	groupAffiliation := dataStruct.Results[0].Connections.GroupAffiliation
	relatives := dataStruct.Results[0].Connections.Relatives
	arrayRelatives := strings.Split(relatives, ",")
	numberRelatives := len(arrayRelatives)
	fmt.Println(arrayRelatives)

	sqlStatement := `
	INSERT INTO superhero (name, fullname, intelligence, power, occupation, image, uuid, group_affiliation, number_relatives)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING uuid`

	uuid := uuid.New().String()

	err = Connection().QueryRow(sqlStatement, name, fullname, intelligence, power, occupation, image, uuid, groupAffiliation, numberRelatives).Scan(&uuid)
	
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON("Uuid inserted:" + uuid)
}

func DeleteSuper (c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	sqlStatement := `
	DELETE FROM superhero
	WHERE uuid = $1;`

	_, err := Connection().Query(sqlStatement, uuid)

	if err != nil {
  	return err
	}

	return c.Status(fiber.StatusOK).JSON(uuid + " deletado")
}
