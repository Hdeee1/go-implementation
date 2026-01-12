package service

import (
	"bytes"
	"encoding/csv"
	"log"
	"strconv"

	"github.com/Hdeee1/go-implementation/internal/model"
)

type ExportService struct {}

func NewExportService() *ExportService {
	return &ExportService{}
}

func (e *ExportService) ExportToCSV(students []*model.Student) ([]byte, error) {
	buf := &bytes.Buffer{}
	write := csv.NewWriter(buf)

	header := []string{"ID", "Name", "Age", "DateOfBirth"}
	if err := write.Write(header); err != nil {
		return nil, err
	}

	progressCh := make(chan int)

	go func() {
		count := len(students)

		for i, student := range students {
			row := []string{
				strconv.Itoa(student.ID),
				student.Name,
				strconv.Itoa(student.Age),
				student.DateOfBirth.Format("2006-01-02"),
			}

			write.Write(row)
			progress := (i + 1) * 100 / count
			progressCh <- progress
		}
		close(progressCh)
	}()

	for p := range progressCh {
		log.Printf("Export progress: %d%%", p)
	}

	write.Flush()

	if err := write.Error(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}