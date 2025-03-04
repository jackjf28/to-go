package entry

import (
	"fmt"
	"testing"
)

func TestReadCsv(t *testing.T) {
	filePath := "dat/test.csv"
	todos, err := ReadCsv(filePath)
	if err != nil {
		t.Errorf("Error parsing csv file %v\n", err)
	}
	for _, todo := range todos {
		fmt.Println(todo.String())
	}
}
