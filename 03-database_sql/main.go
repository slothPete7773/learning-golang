package main

import (
	"database/sql"
	"fmt"

	"example.com/slothpete/database_sql/person"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:my-secret-pw@tcp(localhost)/sys"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// a, err := person.GetPersons(db)
	a, err := person.GetPerson(db, "1")
	if err != nil {
		panic(err)
	}

	fmt.Println(a)

	// newPerson := person.Person{Id: "3", Name: "veer", Score: 10}
	// fmt.Printf("%p\n", &newPerson)

	// err = person.AddPerson(db, newPerson)
	// if err != nil {
	// 	panic(err)
	// }

	// allPeople, err := person.GetPersons(db)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(allPeople)

	// err = person.UpdatePerson(db, newPerson)
	// if err != nil {
	// 	panic(err)
	// }

	err = person.DeletePerson(db, "3")
	if err != nil {
		panic(err)
	}

	// newPerson := person.Person{"2", "veerakit", 10}
	fmt.Println(person.GetPersons(db))

}
