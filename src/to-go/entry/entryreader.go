package entry

import (
	"os"

	"github.com/gocarina/gocsv"
)

func ReadCsv(filePath string) ([]*TodoEntry, error) {
	var todos []*TodoEntry
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return todos, err
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &todos); err != nil {
		return todos, err
	}

	return todos, nil
}
