package main

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"parser/database"
	"parser/files"
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
		db = database.ConnectDB(DBFilename)
		data := files.ReadObjectsCSV(csvFilepath)
		database.BatchInsert(db, data)
		fmt.Println("Dumped csv data into database:", DBFilename)
	} else {
		db = database.ConnectDB(DBFilename)
		changed, added, deleted := files.ParseDiffs(csvPath)
		if len(changed) > 0 {
			database.Update(db, changed)
		}
		if len(added) > 0 {
			database.Create(db, added)
		}
		if len(deleted) > 0 {
			database.Delete(db, deleted)
		}
	}
}
