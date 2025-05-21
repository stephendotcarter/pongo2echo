package pongo2echo

import (
	"embed"
	"io"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
)

func NewRenderer(prefix string, fs *embed.FS) echo.Renderer {
	return Renderer{
		TemplateSet: pongo2.NewSet(prefix, NewLoader(prefix, fs)),
	}
}

// Renderer : Custom Renderer for templates
type Renderer struct {
	trim        string
	TemplateSet *pongo2.TemplateSet
}

// Render : Custom renderer
func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var ctx pongo2.Context
	if data != nil {
		var ok bool
		ctx, ok = data.(pongo2.Context)
		if !ok {
			panic("please pass pongo2.Context to the Render function")
		}
	}
	return pongo2.Must(r.TemplateSet.FromFile(name)).ExecuteWriter(ctx, w)
}
