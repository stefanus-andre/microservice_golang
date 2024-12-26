package repositories

import (
	"database/sql"
	"microservice_golang/user_service/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email) VALUES (?,?)", user.Name, user.Email)
	return err
}

func (r *UserRepository) GetAllUser() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
