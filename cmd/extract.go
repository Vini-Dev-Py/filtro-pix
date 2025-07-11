package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var inputExtract string
var outputExtract string

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extrai relações TXID : EndToEndId de arquivos de log Pix",
	Long: `O comando 'extract' lê um arquivo de log Pix e extrai os pares 
TXID : EndToEndId de cada transação registrada em formato JSON.

Exemplo de uso:
  filtro-pix extract --input pagamentos_pix.log
`,
	Run: func(cmd *cobra.Command, args []string) {
		extractTxidPairs(inputExtract)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
	extractCmd.Flags().StringVarP(&inputExtract, "input", "i", "", "Arquivo de log de entrada")
	extractCmd.Flags().StringVarP(&outputExtract, "output", "o", "", "Arquivo de saída (opcional)")
	extractCmd.MarkFlagRequired("input")
}

// ---------------------
// IMPLEMENTAÇÃO
// ---------------------

type PixData struct {
	Txid       string `json:"txid"`
	EndToEndId string `json:"endToEndId"`
}

type Pagamento struct {
	Txid string    `json:"txid"`
	Pix  []PixData `json:"pix"`
}

func extractTxidPairs(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	var outFile *os.File
	var writer *bufio.Writer

	if outputExtract != "" {
		outFile, err = os.Create(outputExtract)
		if err != nil {
			fmt.Printf("Erro ao criar o arquivo de saída: %v\n", err)
			return
		}
		defer outFile.Close()
		writer = bufio.NewWriter(outFile)
	}

	scanner := bufio.NewScanner(file)
	jsonRegex := regexp.MustCompile(`JSON: (.+)$`)
	contador := 0

	for scanner.Scan() {
		linha := scanner.Text()

		if !strings.Contains(linha, "JSON:") {
			continue
		}

		match := jsonRegex.FindStringSubmatch(linha)
		if len(match) < 2 {
			continue
		}

		var dados Pagamento
		err := json.Unmarshal([]byte(match[1]), &dados)
		if err != nil {
			continue
		}

		for _, p := range dados.Pix {
			linhaFormatada := fmt.Sprintf("%s : %s", p.Txid, p.EndToEndId)

			// Print no console
			fmt.Println(linhaFormatada)

			// Escreve no arquivo, se necessário
			if writer != nil {
				writer.WriteString(linhaFormatada + "\n")
			}
			contador++
		}
	}

	if writer != nil {
		writer.Flush()
		fmt.Printf("\nTotal salvo em '%s': %d pares\n", outputExtract, contador)
	} else {
		fmt.Printf("\nTotal impresso: %d pares\n", contador)
	}
}

