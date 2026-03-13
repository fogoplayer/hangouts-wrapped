package parse

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func ParseChatDirectoryHandle(handle model.ChatDirectoryHandle) {
	groupInfoJson := jsonSchema.GroupInfo_JsonSchema{}
	messagesJson := jsonSchema.Messages_JsonSchema{}

	// TODO parallelize
	parseJson(<-handle.GroupInfo, &groupInfoJson)
	chat := parseGroupInfo(groupInfoJson)
	// fmt.Println(chat.Name)

	parseJson(<-handle.Messages, &messagesJson)

	message := "no messages"
	if len(messagesJson.Messages) > 0 {
		message = messagesJson.Messages[0].Text_
	}
	fmt.Println(chat.Name)
	fmt.Println("\t", message)
}

func ParseUserInfo(bytes []byte) {
	userInfoJson := jsonSchema.UserInfo_JsonSchema{}
	err := parseJson(bytes, &userInfoJson)
	if err != nil {
		fmt.Println("Error parsing user info:", err)
	} else {
		fmt.Println(userInfoJson.User.Name)
	}
}

func parseJson(bytes []byte, destinationPointer any) error {
	if !json.Valid(bytes) {
		return errors.New("invalid json")
	}
	json.Unmarshal(bytes, destinationPointer)
	return nil
}

func parseGroupInfo(groupInfo jsonSchema.GroupInfo_JsonSchema) model.Chat {
	chat := model.Chat{}

	for _, member := range groupInfo.Members {
		chat.Members = append(chat.Members, parseMember(member))
	}

	if groupInfo.Name != "" && groupInfo.Name != "Group Chat" {
		chat.Name = groupInfo.Name
	} else {
		memberNames := []string{}
		for _, member := range chat.Members {
			memberNames = append(memberNames, member.Name)
		}

		chat.Name = "DM with " + strings.Join(memberNames, "/")
	}

	return chat
}

func parseMember(member jsonSchema.GroupInfo_Members_JsonSchema) model.User {
	return model.User{Name: member.Name, Email: member.Email}
}

/////////////
// Message //
/////////////

func parseMessage(message jsonSchema.Message) model.Message {
	return model.Message{
		Creator:   parseCreator(message.Creator),
		TopicId:   message.Topic_Id,
		MessageId: message.Message_Id,

		CreatedDate_:               parseTime(message.Created_Date_),
		Text_:                      message.Text_,
		Message_State_:             parseMessageState(message.Message_State_),
		Deleted_Date_:              parseTime(message.Deleted_Date_),
		Updated_Date_:              parseTime(message.Updated_Date_),
		Quoted_Message_Metadata_:   parseQuotedMessageMetadata(message.Quoted_Message_Metadata_),
		DeletionType_:              parseDeletedMetadata(message.Deletion_Metadata_),
		Annotations_:               parseAnnotations(message.Annotations_),
		Attached_Files_:            parseAttachedFiles(message.Attached_Files_),
		Reactions_:                 parseReactions(message.Reactions_),
		Previous_Message_Versions_: parsePreviousMesssageVersions(message.Previous_Message_Versions_),
	}
}

func parseCreator(creator jsonSchema.Creator) model.Creator {
	return model.Creator{
		Name:        creator.Name,
		Email:       creator.Email,
		UserType_:   creator.User_type_,
		ActingUser_: parseActingUser(creator.Acting_user_),
	}
}

func parseActingUser(actingUser jsonSchema.ActingUser) *model.ActingUser {
	return &model.ActingUser{
		Name:     actingUser.Name,
		UserId:   actingUser.User_id,
		UserType: actingUser.User_type,
	}
}

// 2006-01-02T15:04:05 -070000
var CREATED_DATE = "Monday, January 2 at 3:04:05PM UTC"

func parseTime(datetime string) time.Time {
	time, err := time.Parse(CREATED_DATE, datetime)
	if err != nil {
		fmt.Println("unable to parse time")
		panic(err)
	}
	return time
}

func parseMessageState(messageState string) *model.MessageState {
	if messageState == "DELETED" {
		v := model.DELETED
		return &v
	}
	if messageState == "" {
		return nil
	}
	// else
	panic(fmt.Errorf("Unexpected message state: %s", messageState))
}

func parseQuotedMessageMetadata(quotedMessageMetadata jsonSchema.QuotedMessageMetadata) *model.QuotedMessageMetadata {
	v := model.QuotedMessageMetadata{
		Creator: parseCreator(quotedMessageMetadata.Creator),
		Text:    quotedMessageMetadata.Text,

		AttachedFiles_: parseAttachedFiles(quotedMessageMetadata.Attached_files_),
		Annotations_:   parseAnnotations(quotedMessageMetadata.Annotations_),
	}
	return &v
}

func parseDeletedMetadata(deletedMetadata jsonSchema.DeletionMetadata) *model.DeletionTypeEnum {
	switch model.DeletionType[deletedMetadata.Deletion_type] {
	case model.CREATOR:
		v := model.CREATOR
		return &v
	default:
		fmt.Println(deletedMetadata)
		panic(fmt.Errorf("Probably not actually an error. What does empty deletion metadata look like?"))
	}
}

func parseAnnotations(annotations []jsonSchema.Annotation) []model.Annotation {
	parsedAnnotations := []model.Annotation{}
	for _, annotation := range annotations {
		parsedAnnotations = append(parsedAnnotations, model.Annotation{
			StartIndex:                 annotation.Start_index,
			Length:                     annotation.Length,
			YoutubeMetadata_:           parseYoutubeMetadata(annotation.Youtube_metadata_),
			UrlMetadata_:               parseUrlMetadata(annotation.Url_metadata_),
			VideoCallUrl_:              annotation.Video_call_metadata_.MeetingSpace.MeetingUrl,
			FormatMetadata_:            parseFormatMetadata(annotation.Format_metadata_),
			GsuiteIntegrationMetadata_: parseGsuiteIntegrationMetadata(annotation.Gsuite_integration_metadata_),
			DriveMetadata_:             parseDriveMetadata(annotation.Drive_metadata_),
			InteractionDataUrl_:        parseUrl(annotation.Interaction_data_.Url),
		})
	}
	return parsedAnnotations
}

func parseYoutubeMetadata(youtubeMetadata jsonSchema.YoutubeMetadata) *model.YoutubeMetadata {
	v := model.YoutubeMetadata{
		Id:        youtubeMetadata.Id,
		StartTime: youtubeMetadata.Start_time,
	}
	return &v
}

func parseUrlMetadata(urlMetadata jsonSchema.UrlMetadata) *model.UrlMetadata {
	v := model.UrlMetadata{
		Title:     urlMetadata.Title,
		Snippet:   urlMetadata.Snippet,
		Image_url: urlMetadata.Image_url,
		Url:       parseUrl(urlMetadata.Url),
	}
	return &v
}

func parseUrl(url jsonSchema.Url) string {
	return url.Private_do_not_access_or_else_safe_url_wrapped_value
}

func parseFormatMetadata(formatMetadata jsonSchema.FormatMetadata) *model.FormatMetadata {
	v := model.FormatMetadata{
		FormatType: parseFormatType(formatMetadata.Format_type),
		FontColor_: formatMetadata.Font_color_,
	}
	return &v
}

func parseFormatType(formatType string) model.FormatTypeEnum {
	return model.FormatType[formatType]
}

func parseGsuiteIntegrationMetadata(gsuiteIntegrationMetadata jsonSchema.GsuiteIntegrationMetadata) *model.GsuiteIntegrationMetadata {
	v := model.GsuiteIntegrationMetadata{
		CallData_:          parseCallData(gsuiteIntegrationMetadata.Call_data_),
		CalendarEventData_: parseCalendarEventData(gsuiteIntegrationMetadata.Calendar_event_data_),
		TasksData_:         parseTasksData(gsuiteIntegrationMetadata.Tasks_data_),
	}
	return &v
}

func parseCallData(callData jsonSchema.CallData) *model.CallData {
	v := model.CallData{
		CallStatus: parseCallStatus(callData.Call_status),
	}
	return &v
}

func parseCallStatus(callStatus string) model.CallStatusEnum {
	return model.CallStatus[callStatus]
}

func parseCalendarEventData(calendarEventData jsonSchema.CalendarEventData) *model.CalendarEventData {
	v := model.CalendarEventData{
		Title:     calendarEventData.Calendar_event.Title,
		StartTime: parseTime(calendarEventData.Calendar_event.Start_time.Timed),
		EndTime:   parseTime(calendarEventData.Calendar_event.End_time.Timed),
	}
	return &v
}

func parseTasksData(tasksData jsonSchema.TasksData) *model.TasksData {
	v := model.TasksData{
		Title:          tasksData.Task_properties.Title,
		Completed:      tasksData.Task_properties.Completed,
		Deleted:        tasksData.Task_properties.Deleted,
		Description:    tasksData.Task_properties.Description,
		AssigneeId_:    tasksData.Task_properties.Assignee_.Id,
		OldAssigneeId_: tasksData.Assignee_change_.Old_assignee.Id,
	}
	return &v
}

func parseDriveMetadata(driveMetadata jsonSchema.DriveMetadata) *model.DriveMetadata {
	return &model.DriveMetadata{
		Id:           driveMetadata.Id,
		Title:        driveMetadata.Title,
		ThumbnailUrl: driveMetadata.Thumbnail_url,
	} // TODO apparently &v isn't necessary?
}

func parseAttachedFiles(attachedFiles []jsonSchema.AttachedFile) []model.AttachedFile {
	parsedAttachedFiles := []model.AttachedFile{}
	for _, attachedFile := range attachedFiles {
		parsedAttachedFiles = append(parsedAttachedFiles, model.AttachedFile{
			ExportName:    attachedFile.Export_name,
			OriginalName_: attachedFile.Original_name_,
		})
	}
	return parsedAttachedFiles
}

func parseReactions(reactions []jsonSchema.Reaction) []model.Reaction {
	parsedReactions := []model.Reaction{}
	for _, reaction := range reactions {
		parsedReactions = append(parsedReactions, model.Reaction{
			Emoji:         reaction.Emoji.Unicode,
			ReactorEmails: reaction.Reactor_emails,
		})
	}
	return parsedReactions
}

func parsePreviousMesssageVersions(previousMessageVersions []jsonSchema.PreviousMessageVersion) []model.PreviousMessageVersion {
	parsedPreviousMessageVersions := []model.PreviousMessageVersion{}
	for _, previousMessageVersion := range previousMessageVersions {
		parsedPreviousMessageVersions = append(parsedPreviousMessageVersions, model.PreviousMessageVersion{
			CreatedDate_:            parseTime(previousMessageVersion.Created_date_),
			Text_:                   previousMessageVersion.Text_,
			Updated_date_:           parseTime(previousMessageVersion.Updated_date_),
			Annotations_:            parseAnnotations(previousMessageVersion.Annotations_),
			Attached_files_:         parseAttachedFiles(previousMessageVersion.Attached_files_),
			QuotedMessage_metadata_: parseQuotedMessageMetadata(previousMessageVersion.Quoted_message_metadata_),
		})
	}
	return parsedPreviousMessageVersions
}
