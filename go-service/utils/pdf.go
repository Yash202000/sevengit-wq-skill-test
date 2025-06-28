package utils

import (
	"fmt"
	"go-service/types"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func CreateStudentPDFReport(s types.Student) ([]byte, error) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Student Report", props.Text{
				Top:   3,
				Size:  18,
				Align: consts.Center,
				Style: consts.Bold,
			})
		})
	})

	// Divider
	m.Line(1.0)

	// 2 columns per row
	info := [][2]string{
		{"Name", s.Name},
		{"Email", s.Email},
		{"Class", s.Class},
		{"Section", s.Section},
		{"Roll No", fmt.Sprintf("%d", s.Roll)},
		{"Gender", s.Gender},
		{"Phone", s.Phone},
		{"Father's Name", s.FatherName},
		{"Mother's Name", s.MotherName},
		{"Admission Date", s.AdmissionDate},
		{"Current Address", s.CurrentAddress},
		{"Permanent Address", s.PermanentAddress},
	}

	for i := 0; i < len(info); i += 2 {
		m.Row(8, func() {
			m.Col(6, func() {
				m.Text(fmt.Sprintf("%s: %s", info[i][0], info[i][1]), props.Text{Size: 11})
			})
			if i+1 < len(info) {
				m.Col(6, func() {
					m.Text(fmt.Sprintf("%s: %s", info[i+1][0], info[i+1][1]), props.Text{Size: 11})
				})
			}
		})
	}

	buf, err := m.Output()
	return buf.Bytes(), err
}
