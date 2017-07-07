package liquid

import (
	"bytes"

	"github.com/osteele/liquid/render"
)

// A Template renders a template according to scope.
type Template interface {
	// Render executes the template with the specified bindings.
	Render(Bindings) ([]byte, error)
	// RenderString is a convenience wrapper for Render, that has string input and output.
	RenderString(Bindings) (string, error)
	SetSourcePath(string)
}

type template struct {
	root   render.Node
	config *render.Config
}

// Render executes the template within the bindings environment.
func (t *template) Render(vars Bindings) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := render.Render(t.root, buf, vars, *t.config)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// RenderString is a convenience wrapper for Render, that has string input and output.
func (t *template) RenderString(b Bindings) (string, error) {
	bs, err := t.Render(b)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (t *template) SetSourcePath(filename string) {
	t.config.Filename = filename
}
