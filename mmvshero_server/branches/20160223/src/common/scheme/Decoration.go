package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Decoration struct {
	 Id 	 int32
	 MallId 	 int32
	 Name 	 int32
	 Size 	 string
	 MaxNum 	 int32
	 Icon 	 string
	 DecorationFlashId 	 int32

}

var Decorations []*Decoration
var Decorationmap map[int32]*Decoration

func LoadDecoration(filepath string) {
    fileName := "Decoration.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Decorations)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Decorationmap = make(map[int32]*Decoration)
    for _, v := range Decorations {
        Decorationmap[v.Id] = v
    }
    LogInfo("Load Decoration Scheme Success!")
}
