package model

import (
	"database/sql"
	"errors"
)

type Hero struct {
	id int
	name string
	fullname string
	intelligence string
	power string
	occupation string
	image string
	group_affiliation interface{}
	number_relatives interface{}
	uuid interface{}
}

type Request struct {
	Name string
}

func GetAll () []Hero {
	rows, err := Connection().Query("SELECT * FROM superhero")
	if err != nil {
		panic(err)
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
			panic(err)
		}

		hero := Hero{id, name, fullname, intelligence, power, occupation, image, group_affiliation, number_relatives, uuid}

		heros = append(heros, hero)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return heros
}

func GetByName (name string) (*Hero, error) {
	sqlStatement := `SELECT * FROM superhero WHERE name=$1`

	var super *Hero

	row := Connection().QueryRow(sqlStatement, name)

	err := row.Scan(&super.id, &super.name, &super.fullname, &super.intelligence, &super.power, &super.occupation, &super.image, &super.group_affiliation, &super.number_relatives, &super.uuid)

	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("Super não encontrado")
	case nil:
		return super, nil
	default:
		panic(err)
	}
}