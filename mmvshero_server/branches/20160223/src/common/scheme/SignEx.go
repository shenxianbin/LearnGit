package scheme

var signExmap map[int32]*Sign

func SignProcess() {
	signExmap = make(map[int32]*Sign)
	for _, value := range Signmap {
		index := value.Month<<16 | value.Index
		signExmap[index] = value
	}
}

func SignGet(month int32, index int32) *Sign {
	inx := month<<16 | index
	if v, has := signExmap[inx]; has {
		return v
	}

	return nil
}
