package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type NPC struct {
	 Id 	 int32
	 MapObj 	 int32
	 Name 	 string
	 BaseId 	 int32
	 Lv 	 int32
	 Stage 	 int32
	 Rank 	 int32
	 SkillLv 	 string
	 ItemId 	 int32
	 Drop 	 string
	 Notes 	 string

}

var NPCs []*NPC
var NPCmap map[int32]*NPC

func LoadNPC(filepath string) {
    fileName := "NPC.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &NPCs)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    NPCmap = make(map[int32]*NPC)
    for _, v := range NPCs {
        NPCmap[v.Id] = v
    }
    LogInfo("Load NPC Scheme Success!")
}
