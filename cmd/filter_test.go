package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterCommand(t *testing.T) {
	input := "test_filter_input.log"
	output := "test_filter_output.log"
	content := `Linha normal
Pagamento Pix realizado com sucesso
Outra linha
Erro: Pagamento Pix falhou`

	err := os.WriteFile(input, []byte(content), 0644)
	assert.NoError(t, err)
	defer os.Remove(input)
	defer os.Remove(output)

	inputFilterFile = input
	outputFilterFile = output

	filterCmd.Run(nil, nil)

	data, err := os.ReadFile(output)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "Pagamento Pix")
}
