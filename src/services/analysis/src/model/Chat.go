package model

var ChatsById = make(map[string]Chat)

type Chat struct {
	Name    string
	Members []User
}

type User struct {
	Name  string
	Email string
}
