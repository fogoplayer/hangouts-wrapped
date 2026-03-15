package model

// TODO tree-list-experiments create a general form of this data structure

type Year int
type Month int
type Day int
type Hour int
type Minute int

// type TreeListI[ValueType any] interface {
// 	Values() []ValueType
// }

type YearTreeList map[Year]MonthTreeList

func (yearTreeList YearTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range yearTreeList {
		result = append(result, entry.Values()...)
	}
	return result
}

type MonthTreeList map[Month]DayTreeList

func (monthTreeList MonthTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range monthTreeList {
		result = append(result, entry.Values()...)
	}
	return result
}

type DayTreeList map[Day]HourTreeList

func (dayTreeList DayTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range dayTreeList {
		result = append(result, entry.Values()...)
	}
	return result
}

type HourTreeList map[Hour]MinuteTreeList

func (hourTreeList HourTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range hourTreeList {
		result = append(result, entry.Values()...)
	}
	return result
}

type MinuteTreeList map[Minute]Message

func (minuteTreeList MinuteTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range minuteTreeList {
		result = append(result, entry)
	}
	return result
}

type ChatTreeList map[string]Chat
