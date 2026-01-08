package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Hdeee1/go-implementation/internal/model"
	"github.com/Hdeee1/go-implementation/internal/repository"
)

func GetStudentHandler(repo repository.StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := repo.GetAll()
		if err != nil {
			http.Error(w, "failed to get students", 500)
			return 
		}

		json.NewEncoder(w).Encode(students)
	}
}

func CreateStudentHandler(repo repository.StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.StudentRequest

		json.NewDecoder(r.Body).Decode(&req)
		t, _ := time.Parse("2006-01-02", req.DateOfBirth)
		student := &model.Student{Name: req.Name, Age: req.Age, DateOfBirth: t}

		err := repo.Create(student)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		
		json.NewEncoder(w).Encode(student)
	}
}