package jsonSchema

type UserInfo_JsonSchema struct {
	User                UserInfo_User_JsonSchema             `json:"user"`
	Membership_Info     []UserInfo_MembershipInfo_JsonSchema `json:"membership_info"`
	Starred_Group_Id    []string                             `json:"starred_group_id"`
	Starred_Message_Ids []string                             `json:"starred_message_ids"`
}

type UserInfo_User_JsonSchema struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	User_Type string `json:"user_type"`
}

type UserInfo_MembershipInfo_JsonSchema struct {
	Group_Name       string `json:"group_name"`
	Group_Id         string `json:"group_id"`
	Membership_State string `json:"membership_state"`
}
