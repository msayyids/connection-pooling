package repo

import (
	"context"
	"pgxpool/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	DB *pgxpool.Pool
}

func (r *Repo) AddUser(user model.User) (model.User, error) {
	query := `
        INSERT INTO users (name, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, password
    `

	ctx := context.Background()

	err := r.DB.QueryRow(ctx, query, user.Name, user.Email, user.Password).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
