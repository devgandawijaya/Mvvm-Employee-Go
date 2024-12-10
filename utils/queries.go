package utils

const (
	GetAllUsersQuery = `SELECT * FROM employee`
	CreateUserQuery  = `INSERT INTO employee (name, position, salary) VALUES ($1, $2, $3)`
	UpdateUserQuery  = `UPDATE employee SET name = $1, position = $2, salary = $3 WHERE id = $4`
	DeleteUserQuery  = `DELETE FROM employee WHERE id=$1`
)
