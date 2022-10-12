package models

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Passwd string `json:"passwd,omitempty"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Created int64 `json:"created"`
	Modified int64 `json:"modified"`
}