package db

import (
	"fmt"
	"log"
	pb "nucifera_backend/protos/membership"

	_ "github.com/lib/pq"
)

func (c DBConfig) CheckAuthUserExists(req *pb.Request) (bool, error) {
	var isAuthenticated bool
	log.Println("checking if user exists --------------->")
	sqlStatement := `SELECT CASE WHEN
	 (SELECT COUNT(*) FROM users US WHERE US.email = $1) = 0 THEN 0 ELSE 1 END`
	err := c.DB.QueryRow(sqlStatement, req.Email).Scan(&isAuthenticated)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("User exists = " + fmt.Sprint(isAuthenticated))
	return isAuthenticated, err
}
