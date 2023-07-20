package repository

import (
	"database/sql"
	"errors"

	. "github.com/dwirobbin/edumar-backend/helper"
	. "github.com/dwirobbin/edumar-backend/model/domain"
)

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (repo *AuthRepositoryImpl) FindUser(email, password string) (UserDomain, error) {
	query := `
	SELECT id, username, email, password, loggedin FROM users 
	WHERE email = ? AND password = ?;`

	var user UserDomain
	row := repo.DB.QueryRow(query, email, password)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Loggedin)
	if err != nil {
		return user, errors.New("login failed")
	}

	query = `UPDATE users SET loggedin = true WHERE id = ?;`
	_, err = repo.DB.Exec(query, user.Id)
	PanicIfError(err)

	return user, nil
}

func (repo *AuthRepositoryImpl) FindUsers() ([]UserDomain, error) {
	query := `SELECT id, username, email, loggedin FROM users;`

	rows, err := repo.DB.Query(query)
	PanicIfError(err)
	defer rows.Close()

	var users []UserDomain
	for rows.Next() {
		var user UserDomain
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Loggedin)
		PanicIfError(err)

		users = append(users, user)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	if err := rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (repo *AuthRepositoryImpl) Save(user UserDomain, email string) (UserDomain, error) {
	query := `SELECT id, email FROM users WHERE email = ?;`
	row := repo.DB.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.Email)
	if err == nil {
		return user, errors.New("email already exists")
	}

	query = `
	INSERT INTO users (username, email, password, loggedin, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?);`

	result, err := repo.DB.Exec(query,
		user.Username, email, user.Password, user.Loggedin, user.CreatedAt, user.UpdatedAt,
	)
	PanicIfError(err)

	id, err := result.LastInsertId()
	PanicIfError(err)

	user.Id = uint(id)
	return user, nil
}

func (repo AuthRepositoryImpl) FindUserById(id uint) (UserDomain, error) {
	query := `SELECT id, username, email FROM users WHERE id = ?;`

	var user UserDomain
	row := repo.DB.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
