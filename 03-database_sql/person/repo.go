package person

import (
	"database/sql"
	"errors"
	"fmt"
)

type Person struct {
	Id    string
	Name  string
	Score uint
}

func GetPersons(db *sql.DB) ([]Person, error) {
	if err := db.Ping(); err != nil {
		fmt.Println("Cannot connect to database.")
		return nil, err
	}

	query := "select id, name, score from person;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	// fmt.Println(rows)

	persons := []Person{}
	for rows.Next() {
		person := Person{}
		if err = rows.Scan(&person.Id, &person.Name, &person.Score); err != nil {
			return nil, err
		}

		persons = append(persons, person)

		// fmt.Println(id, name, score)
	}
	return persons, nil
}

func GetPerson(db *sql.DB, id string) (*Person, error) {
	query := "select id, name, score from person where id=?"

	person := Person{}
	row := db.QueryRow(query, id)
	if err := row.Scan(&person.Id, &person.Name, &person.Score); err != nil {
		return nil, err
	}

	return &person, nil
}

func AddPerson(db *sql.DB, person Person) error {
	query := "insert into person (id, name, score) values (?, ?, ?)"
	result, err := db.Exec(query, person.Id, person.Name, person.Score)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 0 {
		return errors.New("No row(s) got created.")
	}

	fmt.Printf("%d row(s) affected.\n", rowsAffected)

	return nil
}

func UpdatePerson(db *sql.DB, person Person) error {
	query := "update person set name=?, score=? where id=?"
	result, err := db.Exec(query, person.Name, person.Score, person.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 0 {
		return errors.New("No row(s) got updated.")
	}

	fmt.Printf("%d row(s) affected.\n", rowsAffected)

	return nil
}

func DeletePerson(db *sql.DB, id string) error {
	query := "delete from person where id=?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 0 {
		return errors.New("No row(s) got deleted.")
	}

	fmt.Printf("%d row(s) affected.\n", rowsAffected)

	return nil
}
