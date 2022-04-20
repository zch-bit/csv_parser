package database

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"parser/models"
)

func ConnectDB(filename string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		fmt.Println("error on open DB", err.Error())
		return nil
	}
	err = db.AutoMigrate(&models.Player{})
	if err != nil {
		fmt.Println("error on migrating DB", err.Error())
		return nil
	}
	return db
}

func BatchInsert(db *gorm.DB, data []*models.Player) {
	start := time.Now()
	db.CreateInBatches(&data, 1000)
	fmt.Println("db batch insertion takes time:", time.Since(start))
}

func Create(db *gorm.DB, data []models.Player) {
	fmt.Printf("+++ db batchinsert: %v\n", data)
	db.CreateInBatches(&data, 10)
}

func Update(db *gorm.DB, data []models.Player) {
	for _, object := range data {
		fmt.Printf("*** db update: %v\n", object)
		db.Save(&object)
	}
}

func Delete(db *gorm.DB, data []models.Player) {
	fmt.Printf("--- db delete: %v\n", data)
	db.Delete(&data)
}
