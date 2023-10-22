package db

import (
	//"database/sql"
	"log"

	_ "github.com/lib/pq"

)

type User struct {
	Email string
}

func (c DBConfig) GetEmails() ([]User, error) {
	sqlStatement := `select u.email 
				from users u`

	rows, err := c.DB.Query(sqlStatement)

	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil

}
