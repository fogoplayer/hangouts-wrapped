package parsed

import "zarinloosli.com/hangouts-wrapped/model"

var ChatsById = make(map[string]Chat)

type Chat struct {
	Name     string
	Members  []User
	Messages YearTreeList
	Id       string
	FSHandle model.FSAgnosticDirectoryHandle
	// TODO Type (DM/Space)
}

type User struct {
	Name  string
	Email string
}
