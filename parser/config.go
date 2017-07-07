package parser

import "github.com/osteele/liquid/expression"

// // Config holds configuration information for parsing and rendering.
type Config struct {
	expression.Config
	Grammar Grammar
}

// NewConfig creates a new Settings.
func NewConfig() Config {
	return Config{}
}
