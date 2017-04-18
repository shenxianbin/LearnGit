package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
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
	 Title 	 int32
	 Detail 	 int32
	 Notes 	 string

}

var NPCs []*NPC
var NPCmap map[int32]*NPC

func LoadNPC(filepath string) {
    fileName := "NPC.json"
	file, err := os.Open(fmt.Sprintf("%s/%s", filepath, fileName))
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	_, err = dec.Token()
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}

	NPCmap = make(map[int32]*NPC)
	for dec.More() {
		var temp NPC
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		NPCs = append(NPCs, &temp)
		NPCmap[temp.Id] = &temp
	}

	LogInfo("Load NPC Scheme Success!")
}
