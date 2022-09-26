package models

type User struct {
	ID int64
	Username string
	Email string
	Passwd string
	FirstName string
	LastName string
	Created int64
	Modified int64
}