package model

type Entry struct {
	ID        string `json:"id"`
	Service   string `json:"service"`
	Password  string `json:"password"`
	Note      string `json:"note"`
	CreatedAt string `json:"created_At"`
}
