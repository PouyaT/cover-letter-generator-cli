package cmd

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/go-pdf/fpdf"
)

var DEFAULT_COVER_LETTER_TEMPLATE = `
Pouya Tavakoli
405-887-7447 | 1pouya.tavakoli@gmail.com
To {{.Company}}

I am interested in the {{.Position}} position at {{.Company}}. I found this position on
your career website and after reading the job description I am confident that I will be a good fit
for the role.

Through my professional career I have learned a lot on managing infrastructure at scale. I
currently work with 50+ clusters per environment with each one having around 100-200 pods.
With the largest cluster having around ~5k pods. I am also part of an on call rotation where I fix
any issues that may arise in production.

On top of the infrastructure management I've developed and maintained Kubernetes controllers and an API that
users utilize to create Kubernetes objects.

I also have experience with handling and getting up to speed with legacy platforms and
spearheading initiatives to migrate customers off of them.

Pouya Tavakoli
`

type CoverLetter struct {
	Company  string
	Position string
}

func generatePDF(content string, filename string) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 12)
	pdf.AddPage()
	pdf.MultiCell(190, 10, content, "", "", false)
	path := "C:/Users/Admin/OneDrive/Documents/Resume and Cover Letter/Cover Letters/" + filename
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	fmt.Printf("Cover Letter generated at %s\n", path)
	return err
}

func GenerateCoverLetter(company string, position string) error {
	data := CoverLetter{
		Company:  company,
		Position: position,
	}
	t := template.Must(template.New("coverLetter").Parse(DEFAULT_COVER_LETTER_TEMPLATE))

	var tplOutput bytes.Buffer
	err := t.Execute(&tplOutput, data)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	outputText := tplOutput.String()
	filename := company + "_Coverletter.pdf"
	err = generatePDF(outputText, filename)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	return err
}
