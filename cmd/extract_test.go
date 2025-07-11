package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractCommand(t *testing.T) {
	input := "test_extract_input.log"
	output := "test_extract_output.txt"
	content := `{"txid":"abc123","endToEndId":"e2e789"}
{"txid":"def456","endToEndId":"e2e000"}
{"foo":"bar"}`

	err := os.WriteFile(input, []byte(content), 0644)
	assert.NoError(t, err)
	defer os.Remove(input)
	defer os.Remove(output)

	inputExtractFile = input
	outputExtractFile = output

	extractCmd.Run(nil, nil)

	data, err := os.ReadFile(output)
	assert.NoError(t, err)
	outputStr := string(data)
	assert.Contains(t, outputStr, "abc123")
	assert.Contains(t, outputStr, "def456")
}
