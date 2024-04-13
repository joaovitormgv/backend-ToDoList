package models

// ToDo struct
type ToDo struct {
	UserId      int    `json:"userId"`
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Hora        string `json:"hora,omitempty"`
	Completed   bool   `json:"completed"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
