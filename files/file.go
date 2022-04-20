package files

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"parser/models"
)

func ReadObjectsCSV(filepath string) []*models.Player {
	start := time.Now()
	file, _ := os.Open(filepath)
	defer file.Close()

	csvReader := csv.NewReader(file)
	objects := make([]*models.Player, 0)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("csv reader error: ", err.Error())
			break
		}
		objects = append(objects, models.ParsePlayer(record))
	}
	fmt.Println("objects has length:", len(objects))
	fmt.Println("read data time:", time.Since(start))
	return objects
}

func GetVersions(path string) (old string, new string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Read dir error,", err.Error())
		return "", ""
	}
	var filenames []string
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}
	old, new = path+"/"+filenames[len(filenames)-2], path+"/"+filenames[len(filenames)-1]
	return
}

func ParseDiffs(path string) ([]models.Player, []models.Player, []models.Player) {
	oldFile, newFile := GetVersions(path)
	fmt.Printf("Processing files: old=%v, new=%v", oldFile, newFile)
	diffs := CalcFileDiff(oldFile, newFile)

	var changed []models.Player
	var added []models.Player
	var removed []models.Player
	for key, val := range diffs {
		if key == "Added" {
			for _, row := range val {
				obj := models.ParsePlayer(row)
				added = append(added, *obj)
			}
		}
		if key == "Removed" {
			for _, row := range val {
				obj := models.ParsePlayer(row)
				removed = append(removed, *obj)
			}
		}
		if key == "Changed" {
			for _, row := range val {
				obj := models.ParsePlayer(row)
				changed = append(changed, *obj)
			}
		}
	}
	fmt.Printf("changed =%v \n", changed)
	fmt.Printf("added =%v \n", added)
	fmt.Printf("removed =%v \n", removed)
	return changed, added, removed
}

func CalcFileDiff(old, new string) map[string][][]string {
	fileOld, _ := os.Open(old)
	defer fileOld.Close()
	fileNew, _ := os.Open(new)
	defer fileNew.Close()
	result, err := csvDiff(fileOld, fileNew)
	if err != nil {
		fmt.Println("err in calculating file diff:", err.Error())
		return nil
	}
	return result
}
