package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcFileDiff(t *testing.T) {
	testDir := "./testdata/"
	x := "123-base-small.csv"
	y := "124-delta-small.csv"
	result := CalcFileDiff(testDir+x, testDir+y)
	assert.Equal(t, result["Changed"][0][0], "69")
	assert.Equal(t, result["Added"][0][0], "24564")
	assert.Equal(t, result["Removed"][0][0], "1615")
}
