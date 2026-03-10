package jsonSchema

type GroupInfo_JsonSchema struct {
	Name    string                         `json:"name"`
	Members []GroupInfo_Members_JsonSchema `json:"members"`
}

type GroupInfo_Members_JsonSchema struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	User_Type string `json:"user_type"`
}
