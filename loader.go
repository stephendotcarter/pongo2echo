package pongo2echo

import (
	`bytes`
	`embed`
	`io`
	`path`

	`github.com/flosch/pongo2`
)

func NewLoader(prefix string, fs *embed.FS) pongo2.TemplateLoader {
	return &Loader{
		prefix: prefix,
		fs:     fs,
	}
}

type Loader struct {
	prefix string
	fs     *embed.FS
}

func (l *Loader) Abs(_, name string) string {
	return path.Join(l.prefix, name)
}

func (l *Loader) Get(path string) (io.Reader, error) {
	b, err := l.fs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
