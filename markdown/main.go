// Package markdown provides a builder pattern for generating Markdown and HTML documents
// with structured formatting, indentation support, and collapsible sections.
//
// The Builder type is the core of this package, offering a fluent API for constructing
// formatted output. It manages indentation levels, line endings, and ensures proper
// formatting of headers, tables, and collapsible detail sections.
//
// Example:
//
//	builder := markdown.NewBuilder()
//	builder.Header("My Document", 1, 0)
//	builder.Header("Section", 2, 0)
//	builder.WriteLine("Some content", 1)
//	output := builder.String()
//
// Builder Methods:
//
// - String(): Get the complete built content
// - Write(): Write raw content (no formatting)
// - WriteLine(): Write a line with indentation
// - WriteEol(): Write end-of-line marker
// - Header(): Create Markdown header (# ## ###)
// - Details(): Create HTML <details> collapsible section with callback-based content
// - HTMLTable(): Create HTML <table> from row iterator
//
// All methods maintain proper indentation and formatting based on depth parameters.
package markdown

import (
	"strings"
)

// Builder constructs Markdown and HTML documents with proper indentation and formatting.
// It uses a fluent API where methods build up content and return the complete result
// via String().
type Builder struct {
	builder      *strings.Builder
	eolString    string
	indentString string
}

// NewBuilder creates a new Builder instance with default settings.
// The builder starts empty and uses two-space indentation with newline line endings.
//
// Example:
//
//	builder := NewBuilder()
//	builder.Header("Title", 1, 0)
//	result := builder.String()
func NewBuilder() *Builder {
	return &Builder{
		builder:      &strings.Builder{},
		eolString:    "\n",
		indentString: "  ",
	}
}
