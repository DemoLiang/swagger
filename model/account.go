package model

// Account example
type Account struct {
	Id   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}
