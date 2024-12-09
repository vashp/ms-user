package domain

type User struct {
	ID    int    `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
}
