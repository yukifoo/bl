package format

import (
	"bytes"
	"testing"

	"github.com/olekukonko/tablewriter"
)

func TestTableFormatterBasic(t *testing.T) {
	// Test that we can create a TableFormatter without errors
	formatter := NewTableFormatter()
	if formatter == nil {
		t.Fatal("NewTableFormatter returned nil")
	}

	// Test that we can set headers without errors
	headers := []string{"Name", "Age", "City"}
	formatter.SetHeader(headers)

	// Test that we can append rows without errors
	row1 := []string{"John", "30", "New York"}
	row2 := []string{"Jane", "25", "London"}
	formatter.AppendRow(row1)
	formatter.AppendRow(row2)

	// Test that render doesn't panic (though we can't easily test output to stdout)
	// In a real test, you might want to redirect stdout, but for now just ensure no panic
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Render() panicked: %v", r)
		}
	}()
	formatter.Render()
}

func TestTableFormatterWithCustomWriter(t *testing.T) {
	// Test with a custom writer to capture output
	var buf bytes.Buffer

	// Create a table with custom writer
	table := tablewriter.NewWriter(&buf)
	formatter := &TableFormatter{table: table}

	// Set headers and add data
	headers := []string{"ID", "Name"}
	formatter.SetHeader(headers)

	row := []string{"1", "Test"}
	formatter.AppendRow(row)

	// Render and check that something was written
	formatter.Render()

	output := buf.String()
	if output == "" {
		t.Fatal("Expected some output, but got empty string")
	}

	// Basic check that our data appears in the output (case insensitive check)
	outputUpper := bytes.ToUpper(buf.Bytes())
	if !bytes.Contains(outputUpper, []byte("ID")) {
		t.Error("Output should contain header 'ID'")
	}
	if !bytes.Contains(outputUpper, []byte("NAME")) {
		t.Error("Output should contain header 'NAME'")
	}
	if !bytes.Contains(outputUpper, []byte("TEST")) {
		t.Error("Output should contain data 'TEST'")
	}
}
