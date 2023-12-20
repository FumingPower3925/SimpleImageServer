package image

import (
	"errors"
	"log"
	"os"

	"github.com/FumingPower3925/SimpleImageServer/pkg/db"
	"github.com/FumingPower3925/SimpleImageServer/pkg/repositories/image/postgres"
)

func GetImageRepository() (IImageRepository, error) {
	switch os.Getenv("DATABASE") {
	case "postgres":
		d := db.New()
		if err := d.DB.Ping(); err != nil {
			log.Fatal(err)
		}
		return postgres.NewPostgresImageRepository(db.New()), nil
	default:
		return nil, errors.New("the requested database is not currently supported")
	}
}
