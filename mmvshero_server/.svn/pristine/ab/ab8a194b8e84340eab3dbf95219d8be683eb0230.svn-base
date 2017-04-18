package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Fb struct {
	 Id 	 int32
	 MapTable 	 string
	 MapPicture 	 string
	 MapSize 	 string
	 NeededLv  	 int32
	 EveryWeek 	 string
	 EveryDay 	 string
	 AttackTimes 	 string
	 LeastCostOrder 	 string
	 ResultCostOrder 	 string
	 LeastRoleExp 	 string
	 ResultRoleExp 	 string
	 RecommendTypeName 	 int32
	 EveryWeekText 	 int32
	 RecommendType 	 int32
	 Effect 	 int32
	 AwardId 	 string
	 Icon 	 string
	 TextIcon 	 string
	 Name 	 int32
	 Details 	 int32

}

var Fbs []*Fb
var Fbmap map[int32]*Fb

func LoadFb(filepath string) {
    fileName := "Fb.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Fbs)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Fbmap = make(map[int32]*Fb)
    for _, v := range Fbs {
        Fbmap[v.Id] = v
    }
    LogInfo("Load Fb Scheme Success!")
}
