package main

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

/*
	export PARSER_NEW_DB=true
*/

func main() {
	DBFilename := "./test.db"
	csvFilepath := "./saved/122-People.csv"
	csvPath := "./saved"

	var db *gorm.DB
	if os.Getenv("PARSER_NEW_DB") == "true" {
		fmt.Println("Set a new database")
		db = ConnectDB(DBFilename)
		data := ReadObjectsCSV(csvFilepath)
		BatchInsert(db, data)
		fmt.Println("Dumped csv data into database:", DBFilename)
	} else {
		db = ConnectDB(DBFilename)
		changed, added, deleted := ParseDiffs(csvPath)
		if len(changed) > 0 {
			Update(db, changed)
		}
		if len(added) > 0 {
			Create(db, added)
		}
		if len(deleted) > 0 {
			Delete(db, deleted)
		}
	}
}
