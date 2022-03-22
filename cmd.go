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

type Column struct {
	Description string `json:"description"`
	Mode        string `json:"mode"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}
type Table []Column

// Main function
func main() {

	if len(os.Args) != 2 {
		log.Fatal("Argument passing failed. Usage: ./cmd comma,separated,paths")
	}

	startDirs := strings.Split(os.Args[1], ",")
	indent := "  "
	var table Table
	tables := make(map[string]Table)
	var paths []string

	// Recursively et all the jsons at the specified location
	for _, startDir := range startDirs {
		paths = append(paths, getJSONs(startDir)...)
	}

	for _, path := range paths {
		table = Table{}
		file, _ := ioutil.ReadFile(path)
		err := json.Unmarshal(file, &table)
		if err != nil {
			log.Fatalf("File reading failed, %v", err)
		}
		tables[path] = table
	}
	descriptionLookup := getDescriptionLookup(tables)
	updatedTablesMode := fillInDefaultMode(tables)
	updatedTables := fillInDescriptions(updatedTablesMode, descriptionLookup)

	writeTables(updatedTables, indent)

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

// Given tables, return map of name -> description
func getDescriptionLookup(tables map[string]Table) map[string]string {
	descriptionLookup := make(map[string]string)
	for _, table := range tables {
		for _, column := range table {
			if column.Description != "" {
				descriptionLookup[column.Name] = column.Description
			}
		}
	}
	return descriptionLookup
}

// Get tables with missing mode, fill in NULLABLE (which is default)
func fillInDefaultMode(tables map[string]Table) map[string]Table {
	for ti, table := range tables {
		updatedMode := false
		for i, column := range table {
			if column.Mode == "" {
				updatedMode = true
				column.Mode = "NULLABLE"
				tables[ti][i] = column
			}
		}
		if !updatedMode {
			delete(tables, ti)
		}
	}
	return tables
}


// Get tables with missing descriptions, try to find & fill them in
// Only return tables that were updated
func fillInDescriptions(tables map[string]Table, descriptionLookup map[string]string) map[string]Table {
	for ti, table := range tables {
		updatedDescription := false
		for i, column := range table {
			if column.Description == "" {
				newDescription := descriptionLookup[column.Name]
				if newDescription != "" {
					updatedDescription = true
					column.Description = newDescription
					tables[ti][i] = column
				}
			}
		}
		if !updatedDescription {
			delete(tables, ti)
		}
	}
	return tables
}

// Writes updated tables to the same locations
func writeTables(tables map[string]Table, indent string) {
	for ti, table := range tables {
		file, _ := json.MarshalIndent(table, "", indent)
		fileInfo, _ := os.Stat(ti)
		_ = ioutil.WriteFile(ti, file, fileInfo.Mode())
	}
}
