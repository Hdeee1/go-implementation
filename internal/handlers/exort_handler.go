package handlers

import (
	"net/http"

	"github.com/Hdeee1/go-implementation/internal/repository"
	"github.com/Hdeee1/go-implementation/internal/service"
)

func ExportStudentHandler(repo repository.StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := repo.GetAll()
		if err != nil {
			http.Error(w, "Failed to get Students", 500)
			return 
		}

		exportService := service.NewExportService()

		csvData, err := exportService.ExportToCSV(students)
		if err != nil {
			http.Error(w, "Failed to export csv", 500)
			return 
		}
		
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=students.csv")
		w.Write(csvData)
	}
}