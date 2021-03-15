package pongo2echo

import (
	`bytes`
	`embed`
	`io`

	`github.com/flosch/pongo2`
)

func NewLoader(fs embed.FS) pongo2.TemplateLoader {
	return &Loader{fs: fs}
}

type Loader struct {
	fs embed.FS
}

func (l *Loader) Abs(_, name string) string {
	return name
}

func (l *Loader) Get(path string) (io.Reader, error) {
	b, err := l.fs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
