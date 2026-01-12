package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/Hdeee1/go-implementation/internal/model"
)

type ImportService struct {}

func NewImportService() *ImportService {
	return &ImportService{}
}

func (i *ImportService) ParseCSV(file io.Reader) ([]model.Student, error) {
	reader := csv.NewReader(file)

	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	studentCh := make(chan *model.Student)
	errCh := make(chan error, 1)

	go func ()  {
		defer close(studentCh)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				errCh <- err
				return 
			}

			id, err := strconv.Atoi(record[0])
			if err != nil {
				errCh <- fmt.Errorf("Data error on line %s: %v", record, err)
				return
			}
			age, _ := strconv.Atoi(record[2])
			dob, _ := time.Parse("2006-01-02", record[3])

			student := &model.Student{
				ID: id,
				Name: record[1],
				Age: age,
				DateOfBirth: dob,
			}

			studentCh <- student
		}
	}()

	var students []model.Student
	for student := range studentCh {
		students = append(students, *student)
	}

	select {
	case err := <- errCh:
		return nil, err
	default:
		return students, nil
	}
}