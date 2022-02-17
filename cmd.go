package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// type Columns struct {
// 	Columns []Column `json:"catlog_nodes"`
// }

type Column struct {
	Name        string `json: "name"`
	Description string `json: "description"`
}

// main function
func main() {

	paths := getJSONs("testjson/")
	print(paths[len(paths)-1])
	file, _ := ioutil.ReadFile("testjson/table1.json")
	var data []Column
	err := json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("File reading failed, %v", err)
	}
	//print(data[len(data)-1].Description)

}

// Given startPath recursively returns paths to all .json files
func getJSONs(startPath string) []string {
	var paths []string
	filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".json") {
			paths = append(paths, path)
		}
		return nil
	})
	return paths
}
