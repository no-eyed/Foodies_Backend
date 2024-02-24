package models

import (
	"foodiesbackend/db"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	ClerkId   string    `json:"clerk_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (clerkid, username, email, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(u.ClerkId, u.Username, u.Email, u.CreatedAt, u.UpdatedAt)
	err = result.Scan(&u.Id)

	return err
}

func (u *User) Update() error {
	query := `UPDATE users SET username = $1, email = $2, updatedAt = $3 WHERE id = $4`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(u.Username, u.Email, u.UpdatedAt, u.Id)
	err = result.Scan(&u.Id)

	return err
}

func (u *User) Delete() error {
	query := `DELETE FROM users WHERE id = $1`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	err = result.Scan(&u.Id)

	return err
}

func GetUserById(id int64) (User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	row := db.DB.QueryRow(query, id)

	var user User

	err := row.Scan(&user.Id, &user.ClerkId, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserIdByClerkid(clerkId string) (int64, error) {
	query := `SELECT id FROM users WHERE clerkid = $1`
	row := db.DB.QueryRow(query, clerkId)

	var id int64

	err := row.Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}
