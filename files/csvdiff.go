package files

import (
	"encoding/csv"
	"fmt"
	"io"
)

// The column index to use as the identifier (matching) column. This should be
// made into an option eventually.
const identifierColumn = 0

// Takes two csv files and starts the diffing operation. Goes through revisionX
// comparing each row against every row in revisionY and then repeats the
// process in reverse.
func csvDiff(fileX, fileY io.Reader) (map[string][][]string, error) {
	revisionXRows, err := csv.NewReader(fileX).ReadAll()
	revisionYRows, err := csv.NewReader(fileY).ReadAll()

	resultMap := map[string][][]string{}

	for _, revisionXRow := range revisionXRows {
		compareRowAgainstRows(revisionXRow, revisionYRows, true, resultMap)
	}

	for _, revisionYRow := range revisionYRows {
		compareRowAgainstRows(revisionYRow, revisionXRows, false, resultMap)
	}

	return resultMap, err
}

// Takes a single row in one revision and compares every row in the other
// revision against that row. If changesAndAdditions is true it only records
// changes and additions, if it is false it only records removals.
func compareRowAgainstRows(row []string, rows [][]string, changesAndAdditions bool, resultMap map[string][][]string) {

	match := false
	// Ideally this algorithm would be more efficient by removing rows once they have been matched.
	for _, comparisonRow := range rows {
		if row[identifierColumn] == comparisonRow[identifierColumn] {
			match = true
			changed := false
			for i := 1; i < len(row); i++ {
				if row[i] != comparisonRow[i] {
					changed = true
				}
			}
			if changed && changesAndAdditions {
				fmt.Println("Changed: ", comparisonRow)
				resultMap["Changed"] = append(resultMap["Changed"], comparisonRow)
			}
		}
	}
	if !match {
		if changesAndAdditions {
			fmt.Println("Removed: ", row)
			resultMap["Removed"] = append(resultMap["Removed"], row)
		} else {
			fmt.Println("Added: ", row)
			resultMap["Added"] = append(resultMap["Added"], row)
		}
	}

}
