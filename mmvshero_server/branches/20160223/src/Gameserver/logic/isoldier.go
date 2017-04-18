package logic

type ISoldier interface {
	GetSchemeId() int32
	GetNum() int32
	GetLevel() int32
	GetStage() int32
	GetSkillLevel() map[int32]int32
	GetExp() int32
	GetTimestamp() int64
	GetEvoSpeedTimeStamp() int64
	GetActive() int32
	GetAutoId() int64
}
