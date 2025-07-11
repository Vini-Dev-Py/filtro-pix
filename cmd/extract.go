package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	inputExtractFile  string
	outputExtractFile string
)

type pixData struct {
	TXID       string `json:"txid"`
	EndToEndID string `json:"endToEndId"`
}

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extrai TXID : EndToEndId de logs JSON Pix",
	Long:  "Lê cada linha JSON de um arquivo e extrai os pares TXID : EndToEndId, exibindo no console ou salvando em arquivo.",
	Run: func(cmd *cobra.Command, args []string) {
		if inputExtractFile == "" {
			fmt.Println("Você precisa fornecer um arquivo de entrada com -i.")
			return
		}

		in, err := os.Open(inputExtractFile)
		if err != nil {
			fmt.Println("Erro ao abrir arquivo:", err)
			return
		}
		defer in.Close()

		var out *os.File
		if outputExtractFile != "" {
			out, err = os.Create(outputExtractFile)
			if err != nil {
				fmt.Println("Erro ao criar arquivo de saída:", err)
				return
			}
			defer out.Close()
		}

		scanner := bufio.NewScanner(in)
		writer := bufio.NewWriter(out)

		count := 0
		for scanner.Scan() {
			line := scanner.Text()
			var data pixData
			if err := json.Unmarshal([]byte(line), &data); err == nil {
				if data.TXID != "" && data.EndToEndID != "" {
					result := fmt.Sprintf("%s : %s", data.TXID, data.EndToEndID)
					if outputExtractFile != "" {
						writer.WriteString(result + "\n")
					} else {
						fmt.Println(result)
					}
					count++
				}
			}
		}

		if outputExtractFile != "" {
			writer.Flush()
			fmt.Printf("✅ %d pares TXID : EndToEndId salvos em %s\n", count, outputExtractFile)
		} else {
			fmt.Printf("✅ %d pares encontrados\n", count)
		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
	extractCmd.Flags().StringVarP(&inputExtractFile, "input", "i", "", "Arquivo de entrada")
	extractCmd.Flags().StringVarP(&outputExtractFile, "output", "o", "", "Arquivo de saída (opcional)")
}
