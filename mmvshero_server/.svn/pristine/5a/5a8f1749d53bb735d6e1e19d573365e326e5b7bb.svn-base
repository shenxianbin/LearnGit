package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func parseExcel(excelFileName, jsonFileName, goFileName string) error {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return err
	}

	type pair map[string]interface{}
	var out []pair
	var keys map[int]string = make(map[int]string)
	var types map[int]string = make(map[int]string)
	var packages map[string]string = make(map[string]string)

	for _, sheet := range xlFile.Sheets {
	RowLoop:
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				for cellIndex, cell := range row.Cells {
					if cell.String() == "" {
						break
					}

					keys[cellIndex] = strings.Trim(cell.String(), " ")
				}
				continue
			}

			if rowIndex == 2 {
				for cellIndex, cell := range row.Cells {
					types[cellIndex] = strings.Trim(cell.String(), " ")
				}
				continue
			}

			if rowIndex < 3 {
				continue
			}

			var p pair = make(map[string]interface{})

			for i := 0; i < len(keys); i++ {
				var cellContent interface{}

				switch types[i] {
				case "string", "package":
					if i >= len(row.Cells) {
						cellContent = ""
						break
					}

					if i == 0 && 0 == strings.Compare(row.Cells[i].String(), "") {
						fmt.Println("Notice, ", excelFileName, " skip empty row", rowIndex)
						break RowLoop
					}

					if types[i] == "package" {
						var data map[string]interface{} = make(map[string]interface{})
						err := json.Unmarshal([]byte(row.Cells[i].Value), &data)
						if err != nil {
							var wrappingArray []interface{}
							err2 := json.Unmarshal([]byte(row.Cells[i].Value), &wrappingArray)
							if err2 != nil {
								return fmt.Errorf("row index %v, row i %v :%s :%s", rowIndex, i, err2, row.Cells[i].String())
							}
							cellContent = wrappingArray
						} else {
							cellContent = data
						}

						temp := parseJsonFields(cellContent, 0)
						if len(temp) > len(packages[keys[i]]) {
							packages[keys[i]] = temp
						}

					} else {
						cellContent = row.Cells[i].Value
					}

				case "int":
					if i >= len(row.Cells) {
						cellContent = 0
						break
					}

					if i == 0 && 0 == strings.Compare(row.Cells[i].String(), "") {
						fmt.Println("Notice, ", excelFileName, " skip empty row", rowIndex)
						break RowLoop
					}

					if row.Cells[i].Value == "" {
						cellContent = 0
						break
					}

					temp, err := row.Cells[i].Float()
					if err != nil {
						fmt.Println("row index:", rowIndex, "cell index:", i, "content:", row.Cells[i].Value, err)
						return err
					}

					cellContent = int64(math.Floor(temp + 0.5))
				default:
					return fmt.Errorf("error type :%v , row index:%v,cell index:%v", types[i], rowIndex, i)
				}

				currentKey := keys[i]
				p[currentKey] = cellContent
			}
			out = append(out, p)
		}
		break
	}

	jsonByte, err := json.MarshalIndent(out, " ", "   ")
	if err != nil {
		return err
	}
	writeFile(jsonFileName, jsonByte)

	_, name := filepath.Split(excelFileName)
	FileName := strings.Replace(name, ".xlsx", "", 1)

	goContent := strings.Replace(template, "{FileName}", FileName, -1)
	goContent = strings.Replace(goContent, "{Fields}", getFields(keys, types, packages), -1)

	err = writeFile(goFileName, []byte(goContent))
	if err != nil {
		return err
	}
	return nil
}

func getFields(keys, types map[int]string, packages map[string]string) string {
	var fields string
	for i := 0; i < len(keys); i++ {
		if types[i] == "int" {
			fields += fmt.Sprintln("\t", keys[i], "\t", "int32")
		} else if types[i] == "string" {
			fields += fmt.Sprintln("\t", keys[i], "\t", "string")
		} else {
			fields += fmt.Sprintln("\t", keys[i], "\t", packages[keys[i]])
		}
	}
	return fields
}

func parseJsonFields(m interface{}, deep int) string {
	str := ""

	switch m.(type) {
	case string, []byte:
		str += "string"
	case int, int32, int64, float64:
		str += "int32"
	case []int, []int32, []int64:
		str += "[]int32"
	case [](interface{}):
		t := m.([](interface{}))
		if len(t) > 0 {
			temp := deep
			str += "[]" + parseJsonFields(t[0], temp+1)
		}
	case map[string]interface{}:
		t := m.(map[string]interface{})
		sorted_keys := make([]string, 0)
		for k, _ := range t {
			sorted_keys = append(sorted_keys, k)
		}
		sort.Strings(sorted_keys)
		str += "*struct {\n"
		for _, k := range sorted_keys {
			temp := deep
			str += outputTab(deep) + fmt.Sprintf("%v %v\n", strings.Title(k), parseJsonFields(t[k], temp+1))
		}

		str += outputTab(deep) + "}"
	}

	return str
}

func outputTab(deep int) string {
	str := ""
	for i := 0; i <= deep; i++ {
		str += fmt.Sprintf("\t")
	}
	return str
}

func writeFile(filename string, data []byte) error {
	var f *os.File
	if _, err := os.Stat(filename); os.IsExist(err) {
		err = os.Remove(filename)
		if err != nil {
			return fmt.Errorf("%s%s", "Remove file failed :", err)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("%s%s", "Create file failed :", err)
	}

	_, err = io.WriteString(f, string(data))
	if err != nil {
		return err
	}
	return nil
}

func getPWD() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}

const template string = `package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type {FileName} struct {
{Fields}
}

var {FileName}s []*{FileName}
var {FileName}map map[int32]*{FileName}

func Load{FileName}(filepath string) {
    fileName := "{FileName}.json"
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

	{FileName}map = make(map[int32]*{FileName})
	for dec.More() {
		var temp {FileName}
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		{FileName}s = append({FileName}s, &temp)
		{FileName}map[temp.Id] = &temp
	}

	LogInfo("Load {FileName} Scheme Success!")
}
`

var goPath string
var dirPth string
var jsonPath string

var inputReader *bufio.Reader

func main() {
	// pwd := "/Users/apple/go/mmvshero_request/trunk/scheme"
	pwd := filepath.Dir(getPWD())

	dirPth = pwd + "/xlsx"
	if _, err := os.Stat(dirPth); os.IsNotExist(err) {
		fmt.Println("excel dir is not exist :", dirPth)
		return
	}

	jsonPath = pwd + "/json"
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		fmt.Println("json dir is not exist :", jsonPath)
		return
	}

	goPath = pwd + "/go"
	if _, err := os.Stat(goPath); os.IsNotExist(err) {
		fmt.Println("json dir is not exist :", goPath)
		return
	}

	suffix := strings.ToUpper(".xlsx")
	filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			jsonName := strings.Replace(filepath.Base(filename), "xlsx", "json", -1)
			goName := strings.Replace(filepath.Base(filename), "xlsx", "go", -1)
			err := parseExcel(filename, jsonPath+"/"+jsonName, goPath+"/"+goName)
			if err != nil {
				fmt.Println(filename, err)
				return err
			}
			fmt.Println(strings.Replace(jsonName, ".json", "", 1), "done.")
		}
		return nil
	})

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please press enter to exit .")
	inputReader.ReadString('\n')
}
