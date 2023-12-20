package postgres

import (
	"context"

	"github.com/FumingPower3925/SimpleImageServer/pkg/db"
	"github.com/FumingPower3925/SimpleImageServer/pkg/models"
)

type postgresUserRepository struct {
	Data *db.Data
}

func NewPostgresUserRepository(data *db.Data) *postgresUserRepository {
	return &postgresUserRepository{Data: data}
}

func (ur *postgresUserRepository) CheckCredentialsUser(ctx context.Context, id string, password string) (bool, error) {
	query := `SELECT * FROM users WHERE id_usuario = $1 AND password = $2`
	row, err := ur.Data.DB.QueryContext(ctx, query, id, password)

	if err != nil {
		return false, err
	}

	return row.Next(), nil
}

func (ur *postgresUserRepository) ExistsCreator(ctx context.Context, id string) (bool, error) {
	query := `SELECT * FROM users WHERE id_usuario = $1`
	row, err := ur.Data.DB.QueryContext(ctx, query, id)

	if err != nil {
		return false, err
	}

	return row.Next(), nil
}

func (ur *postgresUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, password) VALUES($1, $2) RETURNING id;`
	row := ur.Data.DB.QueryRowContext(ctx, query, user.Id, user.Password)
	err := row.Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (ur *postgresUserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id=$1;`
	stmt, err := ur.Data.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
