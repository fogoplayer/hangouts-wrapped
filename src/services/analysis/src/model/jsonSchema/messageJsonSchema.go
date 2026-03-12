package jsonSchema

// TODO remove underscores

// Convention properties with a trailing underscore (_) are optional

type Message struct {
	Creator    Creator `json:"creator"`
	Topic_Id   string  `json:"topic_id"`
	Message_Id string  `json:"message_id"`

	// optional

	Created_Date_              string                   `json:"created_date"`
	Text_                      string                   `json:"text_"`
	Annotations_               []Annotation             `json:"annotations"`
	Attached_Files_            []AttachedFile           `json:"attached_files"`
	Reactions_                 []Reaction               `json:"reactions"`
	Updated_Date_              string                   `json:"updated_date"`
	Previous_Message_Versions_ []PreviousMessageVersion `json:"previous_message_versions"`
	Quoted_Message_Metadata_   QuotedMessageMetadata    `json:"quoted_message_metadata"`
	Message_State_             string                   `json:"message_state"`
	Deleted_Date_              string                   `json:"deleted_date"`
	Deletion_Metadata_         DeletionMetadata         `json:"deletion_metadata"`
}

type Creator struct {
	Name  string `json:"name"`
	Email string `json:"email"`

	// optional

	User_type_   string     `json:"user_type"`
	Acting_user_ ActingUser `json:"acting_user"`
}

type ActingUser struct {
	Name      string `json:"name"`
	User_id   string `json:"user_id"`
	User_type string `json:"user_type"`
	// optional
}

type Annotation struct {
	Start_index int `json:"start_index "`
	Length      int `json:"length"`

	// optional

	Youtube_metadata_            YoutubeMetadata           `json:"youtube_metadata"`
	Url_metadata_                UrlMetadata               `json:"url_metadata"`
	Video_call_metadata_         VideoCallMetadata         `json:"video_call_metadata"`
	Format_metadata_             FormatMetadata            `json:"format_metadata"`
	Gsuite_integration_metadata_ GsuiteIntegrationMetadata `json:"gsuite_integration_metadata"`
	Drive_metadata_              DriveMetadata             `json:"drive_metadata"`
	Interaction_data_            InteractionData           `json:"interaction_data"`
}

type UrlMetadata struct {
	Title     string `json:"title"`
	Snippet   string `json:"snippet"`
	Image_url string `json:"image_url"`
	Url       Url    `json:"url"`
	// optional
}

type Url struct {
	Private_do_not_access_or_else_safe_url_wrapped_value string `json:"private_do_not_access_or_else_safe_url_wrapped_value"`
	// optional
}

type YoutubeMetadata struct {
	Id         string `json:"id"`
	Start_time int    `json:"start_time"`
	// optional
}

type VideoCallMetadata struct {
	MeetingSpace struct {
		MeetingUrl string `json:"meeting_url"`
	} `json:"meeting_space"`
}

type FormatMetadata struct {
	Format_type string `json:"format_type"`

	// optional

	Font_color_ uint `json:"font_color"`
}

// TODO parsed type should return an enum of what it is
type GsuiteIntegrationMetadata struct {
	// optional

	Call_data_           CallData          `json:"call_data"`
	Calendar_event_data_ CalendarEventData `json:"calendar_event_data"`
	Tasks_data_          TasksData         `json:"tasks_data"`
}

type CallData struct {
	Call_status string `json:"call_status"`
	// optional
}

type CalendarEventData struct {
	Calendar_event struct {
		Title      string `json:"title"`
		Start_time struct {
			Timed string `json:"timed"`
			// optional
			// `json:"start_time"`
		}
		End_time struct {
			Timed string `json:"timed"`
		} `json:"end_time"`
		// optional
	} `json:"calendar_event"`
	// optional
}

type TasksData struct {
	Task_properties TaskProperties `json:"task_properties"`

	// optional

	Creation_          struct{} `json:"creation"`
	Completion_change_ struct{} `json:"completion_change"`
	Deletion_change_   struct{} `json:"deletion_change"`
	Assignee_change_   struct {
		Old_assignee struct {
			Id string `json:"id"`
			// optional
		} `json:"old_assignee"`
		// optional
	} `json:"assignee_change"`
}

type TaskProperties struct {
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Deleted     bool   `json:"deleted"`
	Description string `json:"description"`

	// optional

	Assignee_ struct {
		Id string `json:"id"`
		// optional
	} `json:"assignee"`
}

type DriveMetadata struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Thumbnail_url string `json:"thumbnail_url"`
	// optional
}
type InteractionData struct {
	Url Url
	// optional
}

type AttachedFile struct {
	Export_name string `json:"export_name"`

	// optional

	Original_name_ string `json:"original_name"`
}

type Reaction struct {
	Emoji struct {
		Unicode string `json:"unicode"`
		// optional
	} `json:"emoji"`
	Reactor_emails []string `json:"reactor_emails"`
	// optional
}

// OPTIONAL CHECK //

type PreviousMessageVersion struct {
	// optional

	Created_date_ string       `json:"created_date"`
	Text_         string       `json:"text"`
	Updated_date_ string       `json:"updated_date"`
	Annotations_  []Annotation `json:"annotations"`

	// For the following two fields, the computed schema is missing some "optional" fields. This is not a problem; they
	// are optional
	// For the following two fields, the computed schema marks some "optional" fields as "required". This is not a
	// problem; since these are input types rather than output types, considering a required field optional just means
	// that in this case the optional field will always be populated
	// In both cases, the schema was computed from real-world data. Just because none of the ingested examples always did
	// or did not include optional fields does not mean we won't be grateful later for being prepared for those
	// possibilities.

	Attached_files_          []AttachedFile          `json:"attached_files"`
	Quoted_message_metadata_ []QuotedMessageMetadata `json:"quoted_message_metadata"`
}

type QuotedMessageMetadata struct {
	Creator Creator `json:"creator"`
	Text    string  `json:"text"`

	// optional

	Attached_files_ []AttachedFile `json:"attached_files"`
	Annotations_    []Annotation   `json:"annotations"`
}

type DeletionMetadata struct {
	Deletion_type string `json:"deletion_type"`
	// optional
}
