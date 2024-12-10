package repository

import (
    "go-api/model"
    "go-api/utils"
    "github.com/jmoiron/sqlx"
)

type UserRepository struct {
    DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
    var users []model.User
    err := r.DB.Select(&users, utils.GetAllUsersQuery)
    return users, err
}

func (r *UserRepository) CreateUser(user model.User) error {
    _, err := r.DB.Exec(utils.CreateUserQuery, user.Name, user.Position, user.Salary)
    return err
}

func (r *UserRepository) UpdateUser(id int, user model.User) error {
    _, err := r.DB.Exec(utils.UpdateUserQuery, user.Name, user.Position, user.Salary, id)
    return err
}

func (r *UserRepository) DeleteUser(id int) error {
    _, err := r.DB.Exec(utils.DeleteUserQuery, id)
    return err
}
