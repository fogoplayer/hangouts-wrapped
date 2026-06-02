package parsed

// TODO tree-list-experiments create a general form of this data structure

type Year int
type Month int
type Day int
type Hour int
type Minute int

// type TreeListI[ValueType any] interface {
// 	Values() []ValueType
// }

// TODO proper past version handling

type YearTreeList map[Year]MonthTreeList

func (yearTreeList YearTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range yearTreeList {
		result = append(result, entry.Values()...)
	}
	return result
}

func (yearTreeList YearTreeList) Insert(value Message) {
	createdDate := value.CreatedDate_
	if createdDate.IsZero() {
		createdDate = value.Updated_Date_
	}
	year := Year(createdDate.Year())
	if yearTreeList[year] == nil {
		yearTreeList[year] = make(MonthTreeList)
	}
	yearTreeList[year].Insert(value)
}

type MonthTreeList map[Month]DayTreeList

func (monthTreeList MonthTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range monthTreeList {
		result = append(result, entry.Values()...)
	}
	return result
}

func (monthTreeList MonthTreeList) Insert(value Message) {
	createdDate := value.CreatedDate_
	if createdDate.IsZero() {
		createdDate = value.Updated_Date_
	}
	month := Month(createdDate.Month())
	if monthTreeList[month] == nil {
		monthTreeList[month] = make(DayTreeList)
	}
	monthTreeList[month].Insert(value)
}

type DayTreeList map[Day]HourTreeList

func (dayTreeList DayTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range dayTreeList {
		result = append(result, entry.Values()...)
	}
	return result
}

func (dayTreeList DayTreeList) Insert(value Message) {
	createdDate := value.CreatedDate_
	if createdDate.IsZero() {
		createdDate = value.Updated_Date_
	}
	day := Day(createdDate.Day())
	if dayTreeList[day] == nil {
		dayTreeList[day] = make(HourTreeList)
	}
	dayTreeList[day].Insert(value)
}

type HourTreeList map[Hour]MinuteTreeList

func (hourTreeList HourTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range hourTreeList {
		result = append(result, entry.Values()...)
	}
	return result
}

func (hourTreeList HourTreeList) Insert(value Message) {
	createdDate := value.CreatedDate_
	if createdDate.IsZero() {
		createdDate = value.Updated_Date_
	}
	hour := Hour(createdDate.Hour())
	if hourTreeList[hour] == nil {
		hourTreeList[hour] = make(MinuteTreeList)
	}
	hourTreeList[hour].Insert(value)
}

type MinuteTreeList map[Minute][]Message

func (minuteTreeList MinuteTreeList) Values() []Message {
	result := []Message{}
	for _, entry := range minuteTreeList {
		result = append(result, entry...)
	}
	return result
}

// TODO preserve ordering
func (minuteTreeList MinuteTreeList) Insert(value Message) {
	createdDate := value.CreatedDate_
	if createdDate.IsZero() {
		createdDate = value.Updated_Date_
	}
	minute := Minute(createdDate.Minute())
	minuteTreeList[minute] = append(minuteTreeList[minute], value)
}

type ChatTreeList map[string]Chat
