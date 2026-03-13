package model

var ChatsById = make(map[string]Chat)

type Chat struct {
	Name     string
	Members  []User
	Messages []Message // TODO replace with TreeList
}

type User struct {
	Name  string
	Email string
}
