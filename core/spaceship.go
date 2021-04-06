package core

import (
	"ImperialFleet/db"
	"errors"
)

// Spaceship - spaceship data type
type Spaceship struct {
	Id       int        `json:"id"`
	Name     string     `json:"name" validate:"max=100"`
	Class    string     `json:"class" validate:"max=100"`
	Armament []Armament `json:"armament"`
	Crew     int        `json:"crew" validate:"max=999999999"`
	Image    string     `json:"image" validate:"max=100"`
	Value    float64    `json:"value" validate:"max=999999999"`
	Status   string     `json:"status" validate:"max=100"`
}

// ListShips - data type used to return list of spaceships
type ListShips struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Armament struct {
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

// NewSpaceship - Returns a new Spaceship
func NewSpaceship(name string, class string, armament []Armament, crew int, image string, value float64, status string) Spaceship {
	spaceship := Spaceship{}
	spaceship.Name = name
	spaceship.Class = class
	spaceship.Armament = armament
	spaceship.Crew = crew
	spaceship.Image = image
	spaceship.Value = value
	spaceship.Status = status
	return spaceship
}

// CreateSpaceship - stores a new spaceship in the database
func (spaceship *Spaceship) CreateSpaceship() error {
	// prepare our query to prevent sql injection
	stmt, es := db.DB.Prepare("INSERT INTO spaceships (name, class, crew, image, value, status) VALUES(?,?,?,?,?,?)")
	if es != nil {
		return es
	}
	// execute the qury
	exec, er := stmt.Exec(spaceship.Name, spaceship.Class, spaceship.Crew, spaceship.Image, spaceship.Value, spaceship.Status)
	if er != nil {
		return er
	}

	// get our last inserted row ID
	insertID, err := exec.LastInsertId()
	if err != nil {
		return err
	}

	// if spaceship.Armament contains an array of objects
	if len(spaceship.Armament) > 0 {
		sqlStr := "INSERT INTO armaments(spaceship_id, title, qty) VALUES "
		vals := []interface{}{}

		// There may be a better way to insert the map such as a bulk insert
		// but didnt have time to investiage for this test

		// iterate map and build query
		for _, value := range spaceship.Armament {
			sqlStr += "(?, ?, ?),"
			vals = append(vals, insertID, value.Title, value.Qty)
		}

		// trim the last ,
		sqlStr = sqlStr[0 : len(sqlStr)-1]

		// prepare the statement
		stmt, _ = db.DB.Prepare(sqlStr)

		// format all vals at once
		_, err = stmt.Exec(vals...)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get - Returns a new Spaceship from the database
func GetSpaceship(id int) (Spaceship, error) {
	// first get spaceships
	row := db.DB.QueryRow("SELECT * FROM spaceships WHERE id = ?", id)
	spaceship := Spaceship{}
	if err := row.Scan(&spaceship.Id, &spaceship.Name, &spaceship.Class, &spaceship.Crew, &spaceship.Image, &spaceship.Value, &spaceship.Status); err != nil {
		return Spaceship{}, err
	}

	// next get armaments & append them
	rows, _ := db.DB.Query("SELECT title, qty FROM armaments WHERE spaceship_id = ?", id)
	for rows.Next() {
		arnament := Armament{}
		if err := rows.Scan(&arnament.Title, &arnament.Qty); err != nil {
			return Spaceship{}, err
		}
		spaceship.Armament = append(spaceship.Armament, arnament)
	}

	return spaceship, nil
}

// ListSpaceships - Returns list of spaceships from the database
func ListSpaceships(name string, class string, status string) ([]ListShips, error) {

	var spaceshipList []ListShips

	// prepare query
	stmt, err := db.DB.Prepare("SELECT id, name, status FROM spaceships WHERE name = ? AND class = ? AND status = ?")
	if err != nil {
		return []ListShips{}, err
	}

	// execute query
	rows, err := stmt.Query(name, class, status)
	if err != nil {
		return []ListShips{}, err
	}

	for rows.Next() {
		spaceship := ListShips{}
		if err := rows.Scan(&spaceship.Id, &spaceship.Name, &spaceship.Status); err != nil {
			return []ListShips{}, err
		}
		spaceshipList = append(spaceshipList, spaceship)
	}

	return spaceshipList, nil
}

// UpdateSpaceship - updates a spaceship from the database
func UpdateSpaceship(id int, name string, class string, crew int, image string, value float64, status string) error {

	// check exists first
	_, err := GetSpaceship(id)
	if err != nil {
		return errors.New("Spaceship with that ID does not exist")
	}

	stmt, err := db.DB.Prepare("UPDATE spaceships SET name = ?, class = ?, crew = ?, image = ?, value = ?, status = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, class, crew, image, value, status, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSpaceship - deletes a spaceship from the database
func DeleteSpaceship(id int) error {

	// check exists first
	_, err := GetSpaceship(id)
	if err != nil {
		return errors.New("Spaceship with that ID does not exist")
	}

	stmt, err := db.DB.Prepare("DELETE FROM spaceships WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
