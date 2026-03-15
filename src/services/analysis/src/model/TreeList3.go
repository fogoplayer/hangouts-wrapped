package model

type TreeList[KeyType comparable, EntryType, ValueType any] struct {
	Entries map[KeyType]EntryType
}

func (treeList TreeList[KeyType, ValueType, ValueType]) Insert(value ValueType) {

}

var chats = TreeList[ChatName, TreeList[Year, TreeList[Month, TreeList[Day, TreeList[Hour, TreeList[Minute, []Message]]]]]]{}
