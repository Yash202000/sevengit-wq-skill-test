package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-service/types"
	"go-service/utils"
	"io"
	"net/http"
	"os"
)

func FetchStudentFromNodeAPI(studentID string) (types.Student, error) {
	authClient, err := NewAuthClient(os.Getenv("API_BASE_URL")) // No `/api/v1` at end
	if err != nil {
		return types.Student{}, err
	}

	url := fmt.Sprintf("%s/students/%s", authClient.BaseUrl, studentID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Student{}, err
	}

	req.Header.Set("x-csrf-token", authClient.CsrfToken)

	cookieHeader := fmt.Sprintf(
		"accessToken=%s; refreshToken=%s; csrfToken=%s",
		authClient.AccessToken,
		authClient.RefreshToken,
		authClient.CsrfToken,
	)
	req.Header.Set("Cookie", cookieHeader)

	resp, err := authClient.Client.Do(req)
	if err != nil {
		return types.Student{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return types.Student{}, fmt.Errorf("student API returned status %d", resp.StatusCode)
	}

	var student types.Student
	if err := json.Unmarshal(body, &student); err != nil {
		return types.Student{}, errors.New("invalid student response")
	}

	return student, nil
}

func GeneratePDFReportForStudent(studentID string) ([]byte, string, error) {

	student, err := FetchStudentFromNodeAPI(studentID)
	if err != nil {
		return nil, "", err
	}

	pdf, err := utils.CreateStudentPDFReport(student)
	if err != nil {
		return nil, "", err
	}

	return pdf, student.Name, nil
}
