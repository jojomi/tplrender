package tplrender

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestRendering(t *testing.T) {
	assert := assert.New(t)
	FilesystemBackend = afero.NewMemMapFs()
	afs := &afero.Afero{Fs: FilesystemBackend}

	opts := Options{
		TemplateDir:      "templates",
		TemplateFilename: "Test1.tpl",
		OutputDir:        "output",
		OutputFilename:   "result.txt",
	}
	err := afs.WriteFile(filepath.Join(opts.TemplateDir, opts.TemplateFilename), []byte(`My name is {{ .Name }}`), 0640)
	if err != nil {
		assert.FailNow(err.Error())
	}

	Template(opts, map[string]interface{}{
		"Name": "Bob",
	})

	oF := filepath.Join(opts.OutputDir, opts.OutputFilename)
	outputContent, err := afs.ReadFile(oF)
	assert.Nil(err)
	assert.Equal(outputContent, []byte("My name is Bob"))
}

func TestHTMLRendering(t *testing.T) {
	assert := assert.New(t)
	FilesystemBackend = afero.NewMemMapFs()
	afs := &afero.Afero{Fs: FilesystemBackend}

	opts := Options{
		TemplateDir:      "templates",
		TemplateFilename: "Test1.tpl",
		OutputDir:        "output",
		OutputFilename:   "result.html",
	}
	err := afs.WriteFile(filepath.Join(opts.TemplateDir, opts.TemplateFilename), []byte(`<h1>My name is {{ .Name }}</h1>`), 0640)
	if err != nil {
		assert.FailNow(err.Error())
	}

	HTMLTemplate(opts, map[string]interface{}{
		"Name": "Bob",
	})

	oF := filepath.Join(opts.OutputDir, opts.OutputFilename)
	outputContent, err := afs.ReadFile(oF)
	assert.Nil(err)
	assert.Equal(outputContent, []byte("<h1>My name is Bob</h1>"))
}
