package incident_response

import (
	"os"
	"path"
	"path/filepath"
	"text/template"
	"time"

	"github.com/sikalabs/slu/utils/slug_utils"
)

type TemplateVariables struct {
	Title  string
	Date   string
	Author string
	Level  string
}

func CreateIncidentResponseFile(
	prefix string,
	date string,
	title string,
) {
	var err error

	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}

	filename := dateParsed.Format("2006-01-02") + "_" +
		slug_utils.SlugifyUnderscore(title) + ".md"

	tv := TemplateVariables{
		Title:  title,
		Date:   date,
		Author: "---author---",
		Level:  "---level (high|medium|low)---",
	}

	fullPath := path.Join(
		prefix,
		dateParsed.Format("2006/01"),
		filename,
	)
	err = os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	t, err := template.New(fullPath).Parse(IncidentResponseFileTemplate)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(fullPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t.Execute(f, tv)
}
