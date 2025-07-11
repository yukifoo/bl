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
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	
	return &TableFormatter{table: table}
}

// SetHeader sets the table headers
func (t *TableFormatter) SetHeader(headers []string) {
	t.table.SetHeader(headers)
}

// AppendRow appends a row to the table
func (t *TableFormatter) AppendRow(row []string) {
	t.table.Append(row)
}

// Render renders the table to stdout
func (t *TableFormatter) Render() {
	t.table.Render()
}