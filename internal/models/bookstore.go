package models

type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Author struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
