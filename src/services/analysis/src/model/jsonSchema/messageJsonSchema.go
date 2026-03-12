package jsonSchema

// TODO remove underscores

// Convention: properties with a trailing underscore (_) are optional

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
			// optional`json:"start_time"`
		}
		End_time struct {
			Timed string `json:"timed"`
		} `json:"end_time"`
		// optional
	} `json:"calendar_event"`
	// optional
}

type TasksData struct {
	//	tasks_data: {
	//	  required: [task_properties]
	//
	// optional
	//
	//	  creation
	//	  completion_change
	//	  assignee_change
	//	  deletion_change
	//	]

	//          task_properties: {
	//            required: [title completed deleted description]
	//            optional: [assignee]
	//            title: string
	//            completed: boolean
	//            deleted: boolean
	//            description: string
	//            assignee: {
	//              required: [id]

	//              id: string
	//            }
	//          }
	//          creation: { required: [] optional: [] }
	//          completion_change: { required: [] optional: [] }
	//          assignee_change: {
	//            required: [old_assignee]

	//            old_assignee: {
	//              required: [id]

	//              id: string
	//            }
	//          }
	//          deletion_change: { required: [] optional: [] }
	//        }
	//      }
}

type DriveMetadata struct {
	//      drive_metadata: {
	//        required: [id title thumbnail_url]

	//        id: string
	//        title: string
	//        thumbnail_url: string
	//      }
	//      interaction_data: {
	//        required: [url]

	//        url: {
	//          required: [
	//            private_do_not_access_or_else_safe_url_wrapped_value
	//          ]

	//	      private_do_not_access_or_else_safe_url_wrapped_value: string
	//	    }
	//	  }
	//	}
	//
	// ]
}

type AttachedFile struct {
	// attached_files: [
	//
	//	{
	//	  required: [export_name]
	//	  optional: [original_name]
	//	  original_name: string
	//	  export_name: string
	//	}
	//
	// ]
}

type Reaction struct {
	// reactions: [
	//
	//	{
	//	  required: [emoji reactor_emails]
	//	  emoji: {
	//	    required: [unicode]
	//	    unicode: string
	//	  }
	//	  reactor_emails: [string]
	//	}
	//
	// ]
}

type PreviousMessageVersion struct {
	//       previous_message_versions: [
	//         {
	//           required: []
	//           optional: [
	//             text
	//             annotations
	//             updated_date
	//             created_date
	//             attached_files
	//             quoted_message_metadata
	//           ]
	//           created_date: string
	//           text: string
	//           annotations: [
	//             {
	//               required: [start_index length]
	//               optional: [
	//                 gsuite_integration_metadata
	//                 url_metadata
	//                 format_metadata
	//                 youtube_metadata
	//                 drive_metadata
	//               ]
	//               start_index: number
	//               length: number
	//               gsuite_integration_metadata: {
	//                 required: [call_data]

	//                 call_data: {
	//                   required: [call_status]

	//                   call_status: string
	//                 }
	//               }
	//               url_metadata: {
	//                 required: [title snippet image_url url]

	//                 title: string
	//                 snippet: string
	//                 image_url: string
	//                 url: {
	//                   required: [
	//                     private_do_not_access_or_else_safe_url_wrapped_value
	//                   ]

	//                   private_do_not_access_or_else_safe_url_wrapped_value: string
	//                 }
	//               }
	//               format_metadata: {
	//                 required: [format_type]
	//                 optional: [font_color]
	//                 format_type: string
	//                 font_color: number
	//               }
	//               youtube_metadata: {
	//                 required: [id start_time]

	//                 id: string
	//                 start_time: number
	//               }
	//               drive_metadata: {
	//                 required: [id title thumbnail_url]

	//                 id: string
	//                 title: string
	//                 thumbnail_url: string
	//               }
	//             }
	//           ]
	//           updated_date: string
	//           attached_files: [
	//             {
	//               required: [original_name export_name]

	//               original_name: string
	//               export_name: string
	//             }
	//           ]
	//           quoted_message_metadata: {
	//             required: [creator text]
	//             optional: [attached_files annotations]
	//             creator: {
	//               required: [name user_type]
	//               optional: [email]
	//               name: string
	//               email: string
	//               user_type: string
	//             }
	//             text: string
	//             attached_files: [
	//               {
	//                 required: [original_name export_name]

	//                 original_name: string
	//                 export_name: string
	//               }
	//             ]
	//             annotations: [
	//               {
	//                 required: [start_index length]
	//                 optional: [format_metadata youtube_metadata]
	//                 start_index: number
	//                 length: number
	//                 format_metadata: {
	//                   required: [format_type]

	//                   format_type: string
	//                 }
	//                 youtube_metadata: {
	//                   required: [id start_time]

	//	          id: string
	//	          start_time: number
	//	        }
	//	      }
	//	    ]
	//	  }
	//	}
	//
	// ]
}

type QuotedMessageMetadata struct {
	//       quoted_message_metadata: {
	//         required: [creator text]
	//         optional: [attached_files annotations]
	//         creator: {
	//           required: [name user_type]
	//           optional: [acting_user email]
	//           name: string
	//           email: string
	//           user_type: string
	//           acting_user: {
	//             required: [name user_id user_type]

	//             name: string
	//             user_id: string
	//             user_type: string
	//           }
	//         }
	//         text: string
	//         attached_files: [
	//           {
	//             required: [original_name export_name]

	//             original_name: string
	//             export_name: string
	//           }
	//         ]
	//         annotations: [
	//           {
	//             required: [start_index length]
	//             optional: [
	//               format_metadata
	//               url_metadata
	//               gsuite_integration_metadata
	//               youtube_metadata
	//               drive_metadata
	//             ]
	//             start_index: number
	//             length: number
	//             format_metadata: {
	//               required: [format_type]
	//               optional: [font_color]
	//               format_type: string
	//               font_color: number
	//             }
	//             url_metadata: {
	//               required: [title snippet image_url url]

	//               title: string
	//               snippet: string
	//               image_url: string
	//               url: {
	//                 required: [
	//                   private_do_not_access_or_else_safe_url_wrapped_value
	//                 ]

	//                 private_do_not_access_or_else_safe_url_wrapped_value: string
	//               }
	//             }
	//             gsuite_integration_metadata: {
	//               required: [tasks_data]

	//               tasks_data: {
	//                 required: [task_properties creation]

	//                 task_properties: {
	//                   required: [title completed deleted description]

	//                   title: string
	//                   completed: boolean
	//                   deleted: boolean
	//                   description: string
	//                 }
	//                 creation: { required: [] optional: [] }
	//               }
	//             }
	//             youtube_metadata: {
	//               required: [id start_time]

	//               id: string
	//               start_time: number
	//             }
	//             drive_metadata: {
	//               required: [id title thumbnail_url]

	//	        id: string
	//	        title: string
	//	        thumbnail_url: string
	//	      }
	//	    }
	//	  ]
	//	}
}

type DeletionMetadata struct {
	//       deletion_metadata: {
	//         required: [deletion_type]

	//	  deletion_type: string
	//	}
}
