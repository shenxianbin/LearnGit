package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type FbDesertHell struct {
	 Id 	 int32
	 Column1 	 string
	 Column2 	 string
	 Column3 	 string
	 Column4 	 string
	 Column5 	 string
	 Column6 	 string
	 Column7 	 string
	 Column8 	 string
	 Column9 	 string
	 Column10 	 string
	 Column11 	 string
	 Column12 	 string
	 Column13 	 string
	 Column14 	 string
	 Column15 	 string
	 Column16 	 string
	 Column17 	 string
	 Column18 	 string
	 Column19 	 string
	 Column20 	 string
	 Column21 	 string
	 Column22 	 string
	 Column23 	 string
	 Column24 	 string
	 Column25 	 string
	 Column26 	 string
	 Column27 	 string

}

var FbDesertHells []*FbDesertHell
var FbDesertHellmap map[int32]*FbDesertHell

func LoadFbDesertHell(filepath string) {
    fileName := "FbDesertHell.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &FbDesertHells)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    FbDesertHellmap = make(map[int32]*FbDesertHell)
    for _, v := range FbDesertHells {
        FbDesertHellmap[v.Id] = v
    }
    LogInfo("Load FbDesertHell Scheme Success!")
}
