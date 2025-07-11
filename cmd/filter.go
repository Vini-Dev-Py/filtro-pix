package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var inputFilter string
var outputFilter string

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filtra linhas que contenham 'Pagamento Pix'",
	Long: `O comando 'filter' varre um arquivo de log linha por linha e salva 
apenas as linhas que contenham o texto 'Pagamento Pix'.

Exemplo de uso:
  filtro-pix filter --input entrada.log --output pagamentos_pix.log
`,
	Run: func(cmd *cobra.Command, args []string) {
		filterPixLines(inputFilter, outputFilter)
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.Flags().StringVarP(&inputFilter, "input", "i", "", "Arquivo de entrada")
	filterCmd.Flags().StringVarP(&outputFilter, "output", "o", "pagamentos_pix.log", "Arquivo de saída")
	filterCmd.MarkFlagRequired("input")
}

// ---------------------
// IMPLEMENTAÇÃO
// ---------------------

func filterPixLines(inputPath, outputPath string) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo de entrada: %v\n", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo de saída: %v\n", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Pagamento Pix") {
			writer.WriteString(line + "\n")
			count++
		}
	}

	writer.Flush()

	fmt.Printf("Linhas contendo 'Pagamento Pix' salvas em '%s': %d\n", outputPath, count)
}
