package tplrender

import (
	"io"
	"io/ioutil"

	htmlTemplate "html/template"
	textTemplate "text/template"
)

func TemplateReaderWriter(input io.Reader, output io.Writer, data interface{}) error {
	return TemplateReaderWriterNamed("tmpl", input, output, data)
}

func TemplateReaderWriterNamed(templateName string, input io.Reader, output io.Writer, data interface{}) error {
	return TemplateReaderWriterNamedWithFuncMap("tmpl", input, output, textTemplate.FuncMap{}, data)
}

func TemplateReaderWriterNamedWithFuncMap(templateName string, input io.Reader, output io.Writer, funcMap textTemplate.FuncMap, data interface{}) error {
	templateData, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}
	t, err := textTemplate.New(templateName).Funcs(funcMap).Parse(string(templateData))
	if err != nil {
		return err
	}
	err = t.Execute(output, data)
	if err != nil {
		return err
	}
	return nil
}

func HTMLTemplateReaderWriter(input io.Reader, output io.Writer, data interface{}) error {
	return HTMLTemplateReaderWriterNamed("tmpl", input, output, data)
}

func HTMLTemplateReaderWriterNamed(templateName string, input io.Reader, output io.Writer, data interface{}) error {
	return HTMLTemplateReaderWriterNamedWithFuncMap("tmpl", input, output, htmlTemplate.FuncMap{}, data)
}

func HTMLTemplateReaderWriterNamedWithFuncMap(templateName string, input io.Reader, output io.Writer, funcMap htmlTemplate.FuncMap, data interface{}) error {
	templateData, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}
	t, err := htmlTemplate.New(templateName).Funcs(funcMap).Parse(string(templateData))
	if err != nil {
		return err
	}
	err = t.Execute(output, data)
	if err != nil {
		return err
	}
	return nil
}
