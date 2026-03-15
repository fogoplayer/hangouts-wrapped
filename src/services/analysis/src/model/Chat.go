package model

var ChatsById = make(map[string]Chat)

type Chat struct {
	Name     string
	Members  []User
	Messages YearTreeList
}

type User struct {
	Name  string
	Email string
}
