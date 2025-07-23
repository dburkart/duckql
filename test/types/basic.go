package types

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string `ddl:"-"`
}
