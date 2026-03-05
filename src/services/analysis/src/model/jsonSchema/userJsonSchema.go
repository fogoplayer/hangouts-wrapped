package jsonSchema

type UserInfo_JsonSchema struct {
	User                UserInfo_User_JsonSchema
	Membership_Info     []UserInfo_MembershipInfo_JsonSchema
	Starred_Group_Id    []string
	Starred_Message_Ids []string
}

type UserInfo_User_JsonSchema struct {
	Name      string
	Email     string
	User_Type string
}

type UserInfo_MembershipInfo_JsonSchema struct {
	Group_Id         string
	Membership_State string
}
