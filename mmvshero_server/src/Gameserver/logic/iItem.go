package logic

type IItem interface {
	GetSchemeId() int32
	GetType() int32
	GetHeapLimit() int32
	GetLv() int32
	GetUseage() int32
	GetValue() []int32
}
