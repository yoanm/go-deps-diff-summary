package markdown

import (
	"strings"
)

// String returns the complete built Markdown/HTML content as a string.
// It reflects all Write operations that have been performed on the builder.
//
// Returns:
// The complete accumulated content as a string.
func (c *Builder) String() string {
	return c.builder.String()
}

// WriteLine writes a line of content with automatic indentation and line ending.
// It prepends the line with spaces based on indentDepth and appends the end-of-line marker.
//
// Parameters:
//   - line: The content to write
//   - indentDepth: The indentation level (0 = no indent, 1 = 1 indent level, etc.)
//
// Example:
//
//	builder.WriteLine("Content", 0)  // No indent
//	builder.WriteLine("Nested", 1)   // One level of indent (2 spaces)
func (c *Builder) WriteLine(line string, indentDepth int) {
	c.Write(strings.Repeat(c.indentString, indentDepth))
	c.Write(line)
	c.WriteEol()
}

// WriteEol writes an end-of-line marker without any indentation.
// This appends a newline to the output
func (c *Builder) WriteEol() {
	c.Write(c.eolString)
}

// Write appends raw content to the builder without any formatting, indentation, or line endings.
// It's useful for writing content fragments that will be combined with other methods.
//
// Parameters:
//   - v: The raw content string to write
//
// Panics:
// If an error occurs during string concatenation (very rare with strings.Builder).
func (c *Builder) Write(v string) {
	_, err := c.builder.WriteString(v)
	if nil != err {
		panic(err)
	}
}
