package repository

import (
	"database/sql"
	"errors"
	"log"
	"restfull-api/m/v2/domain"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo *UserRepository) GetUserByID(id int) (*domain.User, error) {
	user := &domain.User{}
	row := repo.Db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user *domain.User) error {
	err := repo.Db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	return err
}

func (repo *UserRepository) UpdateUser(user *domain.User) (sql.Result, error) {
	res, err := repo.Db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
	return res, err
}

func (repo *UserRepository) DeleteUser(id int) (sql.Result, error) {
	res, err := repo.Db.Exec("DELETE FROM users WHERE id = $1", id)
	return res, err
}

func (repo *UserRepository) ListUsers() ([]*domain.User, error) {
	rows, err := repo.Db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
