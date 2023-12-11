package main

import (
	"os"
)

func GenerateTemplate() (string, error) {
	templateFileName := "input.html"
	templateFilePath := "input.html"

	type structure struct {
		CheckboxState bool
	}

	object := structure{
		CheckboxState: true,
	}

	body, err := ParseTemplateFile(templateFileName, templateFilePath, object)
	if err != nil {
		return "Failure!", err
	}
	buff, err := GeneratePDF(body)
	if err != nil {
		return "Failure!", err
	}

	pdfFilePath := "output.pdf"

	err = os.WriteFile(pdfFilePath, buff.Bytes(), 0644)
	if err != nil {
		return "Failure!", err
	}

	return "Success!", nil
}
