package models

// ToDo struct
type ToDo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
