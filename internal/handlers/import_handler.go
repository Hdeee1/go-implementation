package handlers

import (
	"net/http"

	"github.com/Hdeee1/go-implementation/internal/repository"
	"github.com/Hdeee1/go-implementation/internal/service"
)

func ImportStudentHandler(repo repository.StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Failed to parse form", 400)
			return 
		}

		file, _, err := r.FormFile("File")
		if err != nil {
			http.Error(w, "Failed to get file", 400)
			return 
		}
		defer file.Close()

		importService := service.NewImportService()
		students, err := importService.ParseCSV(file)
		if err != nil {
			http.Error(w, "Failed to parse CSV: "+err.Error(), 500)
			return 
		}

		for _, student := range students {
			err := repo.Create(&student)
			if err != nil {
				http.Error(w, "Failed to insert student: "+err.Error(), 500)
				return 
			}
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Successfully imported students"))
	}
}