package model

// TODO tree-list-experiments create a general form of this data structure

type Year int
type Month int
type Day int
type Hour int
type Minute int

type TreeListI[ValueType any] interface {
	Values() []ValueType
}

type YearTreeList map[Year]MonthTreeList

type MonthTreeList map[Month]DayTreeList
type DayTreeList map[Day]HourTreeList
type HourTreeList map[Hour]MinuteTreeList
type MinuteTreeList map[Minute]Message

type ChatTreeList map[string]Chat
