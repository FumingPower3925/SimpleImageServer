package user

import (
	"errors"
	"log"
	"os"

	"github.com/FumingPower3925/SimpleImageServer/pkg/db"
	"github.com/FumingPower3925/SimpleImageServer/pkg/repositories/user/postgres"
)

func GetUserRepository() (IUserRepository, error) {
	switch os.Getenv("DATABASE") {
	case "postgres":
		d := db.New()
		if err := d.DB.Ping(); err != nil {
			log.Fatal(err)
		}
		return postgres.NewPostgresUserRepository(db.New()), nil
	default:
		return nil, errors.New("the requested database is not currently supported")
	}
}
