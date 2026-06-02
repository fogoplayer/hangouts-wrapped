package parsed

var ChatsById = make(map[string]Chat)

type Chat struct {
	Name     string
	Members  []User
	Messages YearTreeList
	Id       string
	// TODO Type (DM/Space)
}

type User struct {
	Name  string
	Email string
}
