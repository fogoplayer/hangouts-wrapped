package model

import (
	. "time"
)

type Message struct {
	Creator    Creator
	Topic_Id   string
	Message_Id string

	// optional

	Created_Date_              Time
	Text_                      string
	Message_State_             *MessageState
	Deleted_Date_              Time
	Updated_Date_              Time
	Quoted_Message_Metadata_   *QuotedMessageMetadata
	DeletionType_              *DeletionTypeEnum
	Annotations_               []Annotation
	Attached_Files_            []AttachedFile
	Reactions_                 []Reaction
	Previous_Message_Versions_ []PreviousMessageVersion
}

// ///////////////// //
// MessageState Enum //
// ///////////////// //

type MessageState int

const (
	DELETED MessageState = iota
)

type Creator struct {
	Name  string
	Email string

	// optional

	UserType_   string
	ActingUser_ *ActingUser
}

type ActingUser struct {
	Name     string
	UserId   string
	UserType string
	// optional
}

type Annotation struct {
	StartIndex int
	Length     int

	// optional

	YoutubeMetadata_           *YoutubeMetadata
	UrlMetadata_               *UrlMetadata
	VideoCallMetadata_         *VideoCallMetadata
	FormatMetadata_            *FormatMetadata
	GsuiteIntegrationMetadata_ *GsuiteIntegrationMetadata
	DriveMetadata_             *DriveMetadata
	InteractionDataUrl_        string
}

type UrlMetadata struct {
	Title     string
	Snippet   string
	Image_url string
	Url       string
	// optional
}

type YoutubeMetadata struct {
	Id        string
	StartTime int
	// optional
}

type VideoCallMetadata struct {
	MeetingUrl string
}

type FormatMetadata struct {
	FormatType string

	// optional

	FontColor_ uint
}

// TODO make sure to do an exhaustiveness check in the converter
type FormatTypeEnum int

const (
	BOLD FormatTypeEnum = iota
	HIDDEN
	ITALIC
	BULLETED_LIST
	BULLETED_LIST_ITEM
	STRIKE
)

var FormatType = map[string]FormatTypeEnum{
	"BOLD":               BOLD,
	"ITALIC":             ITALIC,
	"HIDDEN":             HIDDEN,
	"BULLETED_LIST":      BULLETED_LIST,
	"BULLETED_LIST_ITEM": BULLETED_LIST_ITEM,
	"STRIKE":             STRIKE,
}

// TODO parsed type should return an enum of what it is
type GsuiteIntegrationMetadata struct {
	// optional

	CallData_          *CallData
	CalendarEventData_ *CalendarEventData
	TasksData_         *TasksData
}

type CallData struct {
	CallStatus string
	// optional
}

type CalendarEventData struct {
	Title     string
	StartTime Time
	EndTime   Time
}

type TasksData struct {
	Title       string
	Completed   bool
	Deleted     bool
	Description string

	// optional

	AssigneeId_    string
	OldAssigneeId_ string
}

type DriveMetadata struct {
	Id           string
	Title        string
	ThumbnailUrl string
	// optional
}

type AttachedFile struct {
	ExportName string

	// optional

	OriginalName_ string
}

type Reaction struct {
	Emoji         string
	ReactorEmails []string
	// optional
}

type PreviousMessageVersion struct {
	// optional

	CreatedDate_  Time
	Text_         string
	Updated_date_ Time
	Annotations_  []Annotation

	// For the following two fields, the computed schema is missing some "optional" fields. This is not a problem; they
	// are optional
	// For the following two fields, the computed schema marks some "optional" fields as "required". This is not a
	// problem; since these are input types rather than output types, considering a required field optional just means
	// that in this case the optional field will always be populated
	// In both cases, the schema was computed from real-world data. Just because none of the ingested examples always did
	// or did not include optional fields does not mean we won't be grateful later for being prepared for those
	// possibilities.

	Attached_files_         []AttachedFile
	QuotedMessage_metadata_ []QuotedMessageMetadata
}

type QuotedMessageMetadata struct {
	Creator Creator
	Text    string

	// optional

	AttachedFiles_ []AttachedFile
	Annotations_   []Annotation
}

type DeletionTypeEnum int

const (
	CREATOR DeletionTypeEnum = iota
)

var DeletionType = map[string]DeletionTypeEnum{
	"CREATOR": CREATOR,
}
