package tplrender

import (
	htmlTemplate "html/template"
	"io"
	"os"
	"path/filepath"
	textTemplate "text/template"

	"github.com/spf13/afero"
)

// FilesystemBackend allows to specifiy another filesystem backend, default is the one from the os package
var FilesystemBackend = afero.NewOsFs()

type Options struct {
	TemplateDir      string
	TemplateFilename string

	OutputDir      string
	OutputFilename string

	NoCreateOutputDir bool
}

func Template(opts Options, data interface{}) error {
	return TemplateWithFuncMap(opts, textTemplate.FuncMap{}, data)
}

func TemplateWithFuncMap(opts Options, funcMap textTemplate.FuncMap, data interface{}) error {
	r, rDeferFunc, err := getReader(opts)
	if err != nil {
		return err
	}
	defer rDeferFunc()

	w, wDeferFunc, err := getWriter(opts)
	if err != nil {
		return err
	}
	defer wDeferFunc()

	err = TemplateReaderWriter(r, w, data)
	return err
}

func HTMLTemplate(opts Options, data interface{}) error {
	return HTMLTemplateWithFuncMap(opts, htmlTemplate.FuncMap{}, data)
}

func HTMLTemplateWithFuncMap(opts Options, funcMap htmlTemplate.FuncMap, data interface{}) error {
	r, rDeferFunc, err := getReader(opts)
	if err != nil {
		return err
	}
	defer rDeferFunc()

	w, wDeferFunc, err := getWriter(opts)
	if err != nil {
		return err
	}
	defer wDeferFunc()

	err = HTMLTemplateReaderWriter(r, w, data)
	return err
}

func getReader(opts Options) (io.Reader, func(), error) {
	// read file
	inputFilename := filepath.Join(opts.TemplateDir, opts.TemplateFilename)

	// execute template
	r, err := FilesystemBackend.Open(inputFilename)
	deferFunc := func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}
	return r, deferFunc, err
}

func getWriter(opts Options) (io.Writer, func(), error) {
	// ensure output dir
	if !opts.NoCreateOutputDir {
		if _, err := FilesystemBackend.Stat(opts.OutputDir); os.IsNotExist(err) {
			FilesystemBackend.MkdirAll(filepath.Dir(opts.OutputFilename), 0750)
		}
	}

	// execute template and thus write output file
	outputFilename := filepath.Join(opts.OutputDir, opts.OutputFilename)
	w, err := FilesystemBackend.Create(outputFilename)
	deferFunc := func() {
		if err := w.Close(); err != nil {
			panic(err)
		}
	}
	return w, deferFunc, err
}
