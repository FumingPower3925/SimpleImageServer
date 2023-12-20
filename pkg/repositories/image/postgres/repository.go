package postgres

import (
	"context"
	"time"

	"github.com/FumingPower3925/SimpleImageServer/pkg/db"
	"github.com/FumingPower3925/SimpleImageServer/pkg/models"
)

type postgresImageRepository struct {
	Data *db.Data
}

func NewPostgresImageRepository(data *db.Data) *postgresImageRepository {
	return &postgresImageRepository{Data: data}
}

func (ir *postgresImageRepository) CreateImage(ctx context.Context, image *models.ControllersImage) error {
	query := `INSERT INTO images (title, description, keywords, author, creator, capture_date, storage_date, filename) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING title, description, keywords, author, creator, capture_date, filename;`
	row := ir.Data.DB.QueryRowContext(ctx, query, image.Title, image.Description, image.Keywords, image.Author,
		image.Creator, image.CaptureDate, time.Now(), image.Filename)
	err := row.Scan(&image.Title, &image.Description, &image.Keywords, &image.Author, &image.Creator, &image.CaptureDate, &image.Filename)
	if err != nil {
		return err
	}
	return nil
}

func (ir *postgresImageRepository) GetAllImages(ctx context.Context) (*[]models.RepositoryImage, error) {
	query := `SELECT * FROM images`
	rows, err := ir.Data.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var images []models.RepositoryImage
	for rows.Next() {
		var image models.RepositoryImage
		rows.Scan(&image.Title, &image.Description, &image.Keywords, &image.Author, &image.Creator, &image.CaptureDate, &image.Filename)
		images = append(images, image)
	}
	return &images, nil
}

func (ir *postgresImageRepository) GetByAttributeImage(ctx context.Context, image *models.ControllersImage) (*[]models.RepositoryImage, error) {
	query := `SELECT * FROM images`
	c := 0
	if image.Title != "" {
		query = query + ` WHERE title = $1`
		c++
	} else if !(image.Author == "" && image.Creator == "" && image.Keywords == "" && image.CaptureDate == "") {
		query = query + `WHERE`
	}
	if image.Keywords != "" {
		if image.Title != "" {
			query = query + `AND `
		}
		query = query + `keywords = $2`
		c++
	}
	if image.Author != "" {
		if image.Title != "" || image.Keywords != "" {
			query = query + `AND `
		}
		query = query + `author = $3`
		c++
	}
	if image.Creator != "" {
		if image.Title != "" || image.Keywords != "" || image.Author != "" {
			query = query + `AND `
		}
		query = query + `creator = $4`
		c++
	}
	if image.CaptureDate != "" {
		if image.Title != "" || image.Keywords != "" || image.Author != "" || image.CaptureDate != "" {
			query = query + `AND `
		}
		query = query + `capture_date = $5`
		c++
	}
	rows, err := ir.Data.DB.QueryContext(ctx, query, image.Title, image.Keywords, image.Author, image.Creator, image.CaptureDate)
	if err != nil {
		return nil, err
	}
	var images []models.RepositoryImage
	for rows.Next() {
		var image models.RepositoryImage
		rows.Scan(&image.Title, &image.Description, &image.Keywords, &image.Author, &image.Creator, &image.CaptureDate, &image)
		images = append(images, image)
	}
	return &images, nil
}

func (ir *postgresImageRepository) UpdateImage(ctx context.Context, filename string, image *models.ControllersImage) error {
	query := `UPDATE images set title=$1, description=$2, author=$3, keywords=$4, capture_date=$5 WHERE filename=$6;`
	stmt, err := ir.Data.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, image.Title, image.Description, image.Author, image.Keywords, image.CaptureDate, filename)
	if err != nil {
		return err
	}
	return nil
}

func (ir *postgresImageRepository) DeleteImage(ctx context.Context, filename string) error {
	query := `DELETE FROM images WHERE filename=$1;`
	stmt, err := ir.Data.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, filename)
	if err != nil {
		return err
	}

	return nil
}
