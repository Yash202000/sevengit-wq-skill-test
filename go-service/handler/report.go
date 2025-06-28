package handler

import (
	"fmt"
	"net/http"

	"go-service/service"

	"github.com/gorilla/mux"
)

func GenerateStudentsReportHandler(w http.ResponseWriter, r *http.Request) {
	studentID := mux.Vars(r)["id"]

	pdfBytes, studentName, err := service.GeneratePDFReportForStudent(studentID)
	if err != nil {
		http.Error(w, "Failed to generate PDF: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", studentName))
	w.Write(pdfBytes)
}
