package models

type User struct {
	ID    int    `json:"id" db:"id" `
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Posts struct {
	ID          int    `json:"id" db:"id" `
	Name        string `json:"name"`
	Description string `json:"description"`
}
