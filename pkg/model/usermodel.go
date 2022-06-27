package model

type User struct {
	Id                int
	Email             string
	Username          string
	EncryptedPassword []byte
}
