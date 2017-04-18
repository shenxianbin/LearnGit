package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type HeroCreatePlan struct {
	 Id 	 int32
	 DeathTime 	 int32
	 ReliveDeathTime 	 int32
	 Time1 	 int32
	 Time2 	 int32
	 Time3 	 int32
	 Time4 	 int32
	 Time5 	 int32
	 Time6 	 int32
	 Time7 	 int32
	 Time8 	 int32
	 Time9 	 int32
	 Time10 	 int32
	 Time11 	 int32
	 Time12 	 int32
	 Time13 	 int32
	 Time14 	 int32
	 Time15 	 int32
	 Time16 	 int32
	 Time17 	 int32
	 Time18 	 int32
	 Time19 	 int32
	 Time20 	 int32
	 Time21 	 int32
	 Time22 	 int32
	 Time23 	 int32
	 Time24 	 int32
	 Time25 	 int32
	 Time26 	 int32
	 Time27 	 int32
	 Time28 	 int32
	 Time29 	 int32
	 Time30 	 int32
	 Time31 	 int32
	 Time32 	 int32
	 Time33 	 int32
	 Time34 	 int32
	 Time35 	 int32
	 Time36 	 int32
	 Time37 	 int32
	 Time38 	 int32
	 Time39 	 int32
	 Time40 	 int32
	 Time41 	 int32
	 Time42 	 int32
	 Time43 	 int32
	 Time44 	 int32
	 Time45 	 int32
	 Time46 	 int32
	 Time47 	 int32
	 Time48 	 int32
	 Time49 	 int32
	 Time50 	 int32
	 Time51 	 int32
	 Time52 	 int32
	 Time53 	 int32
	 Time54 	 int32
	 Time55 	 int32
	 Time56 	 int32
	 Time57 	 int32
	 Time58 	 int32
	 Time59 	 int32
	 Time60 	 int32

}

var HeroCreatePlans []*HeroCreatePlan
var HeroCreatePlanmap map[int32]*HeroCreatePlan

func LoadHeroCreatePlan(filepath string) {
    fileName := "HeroCreatePlan.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &HeroCreatePlans)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    HeroCreatePlanmap = make(map[int32]*HeroCreatePlan)
    for _, v := range HeroCreatePlans {
        HeroCreatePlanmap[v.Id] = v
    }
    LogInfo("Load HeroCreatePlan Scheme Success!")
}
