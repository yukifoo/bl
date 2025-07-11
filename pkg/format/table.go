// Package format provides formatting utilities for CLI output
package format

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// TableFormatter provides table formatting functionality
type TableFormatter struct {
	table *tablewriter.Table
}

// NewTableFormatter creates a new table formatter
func NewTableFormatter() *TableFormatter {
	// Use NewWriter which provides a basic table without complex configuration
	table := tablewriter.NewWriter(os.Stdout)

	return &TableFormatter{table: table}
}

// SetHeader sets the table headers
func (t *TableFormatter) SetHeader(headers []string) {
	// Convert []string to []any for the new API
	headerAny := make([]any, len(headers))
	for i, h := range headers {
		headerAny[i] = h
	}
	t.table.Header(headerAny...)
}

// AppendRow appends a row to the table
func (t *TableFormatter) AppendRow(row []string) {
	// Convert []string to []any for the new API
	rowAny := make([]any, len(row))
	for i, r := range row {
		rowAny[i] = r
	}
	_ = t.table.Append(rowAny...)
}

// Render renders the table to stdout
func (t *TableFormatter) Render() {
	_ = t.table.Render()
}
