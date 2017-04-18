package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type FbDesertNormal struct {
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

var FbDesertNormals []*FbDesertNormal
var FbDesertNormalmap map[int32]*FbDesertNormal

func LoadFbDesertNormal(filepath string) {
    fileName := "FbDesertNormal.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &FbDesertNormals)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    FbDesertNormalmap = make(map[int32]*FbDesertNormal)
    for _, v := range FbDesertNormals {
        FbDesertNormalmap[v.Id] = v
    }
    LogInfo("Load FbDesertNormal Scheme Success!")
}
