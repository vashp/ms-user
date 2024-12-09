package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/vashp/ms-user/src/cmd/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query, args, err := squirrel.Insert("users").
		Columns("email").
		Values(user.Email).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query, args, err := squirrel.Select("id", "email").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user domain.User
	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&user.ID, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query, args, err := squirrel.Update("users").
		Set("email", user.Email).
		Where(squirrel.Eq{"id": user.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query, args, err := squirrel.Delete("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
