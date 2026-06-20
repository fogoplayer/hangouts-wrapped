package state

import (
	"time"

	"zarinloosli.com/hangouts-wrapped/model/parsed"
)

var _ = func() struct{} {
	AllChats.OnChange(allChatsListener)
	return struct{}{}
}()

var AllChats = SetApplicationState[*parsed.Chat]{}
var allChatsListener = func(changed *parsed.Chat) {
	if AllChats.Includes(changed) {
		IncludedChatsFilter.Add(changed)
	} else {
		IncludedChatsFilter.Delete(changed)
	}
}

var MinDateFilter = ApplicationState[time.Time]{
	value: time.Time{},
}

var MaxDateFilter = ApplicationState[time.Time]{
	value: time.Now(),
}

var IncludedUsersFilter = SetApplicationState[string]{}

var IncludedChatsFilter = SetApplicationState[*parsed.Chat]{}

func CountMessagesByYear() map[time.Time]int {
	chats := GetFilteredChats()

	countsByYear := make(map[time.Time]int)

	for _, chat := range chats {
		for year, monthTreeList := range chat.Messages {
			countsByYear[time.Date(int(year), 1, 1, 0, 0, 0, 0, time.UTC)] += len(monthTreeList.Values())
		}
	}

	return countsByYear
}

func CountMessagesByMonthAndYear() map[time.Time]int {
	countsByMonth := make(map[time.Time]int)

	for _, chat := range GetFilteredChats() {
		for year, monthTreeList := range chat.Messages {
			for month, dayTreeList := range monthTreeList {
				countsByMonth[time.Date(int(year), time.Month(month), 1, 0, 0, 0, 0, time.UTC)] += len(dayTreeList.Values())
			}
		}
	}

	return countsByMonth
}

func CountMessagesByHour() []int {
	countsByHour := make([]int, 24)

	for _, chat := range GetFilteredChats() {
		for _, monthTreeList := range chat.Messages {
			for _, dayTreeList := range monthTreeList {
				for _, hourTreeList := range dayTreeList {
					for hour, minuteTreeList := range hourTreeList {
						countsByHour[hour] += len(minuteTreeList.Values())
					}
				}
			}
		}
	}
	return countsByHour
}

func CountMessagesByUser() map[string]int {
	chats := GetFilteredChats()
	messagesByUser := make(map[string]int)
	for _, chat := range chats {
		for _, message := range chat.Messages.Values() {
			messagesByUser[message.Creator.String()] += 1
		}
	}

	return messagesByUser
}
