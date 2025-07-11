package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	inputFilterFile  string
	outputFilterFile string
)

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filtra linhas com 'Pagamento Pix' de um arquivo de log",
	Long:  "Lê um arquivo linha por linha e salva apenas as que contêm 'Pagamento Pix' em um novo arquivo.",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFilterFile == "" || outputFilterFile == "" {
			fmt.Println("Você precisa fornecer os arquivos de entrada e saída usando os flags -i e -o.")
			return
		}

		in, err := os.Open(inputFilterFile)
		if err != nil {
			fmt.Println("Erro ao abrir arquivo de entrada:", err)
			return
		}
		defer in.Close()

		out, err := os.Create(outputFilterFile)
		if err != nil {
			fmt.Println("Erro ao criar arquivo de saída:", err)
			return
		}
		defer out.Close()

		scanner := bufio.NewScanner(in)
		writer := bufio.NewWriter(out)

		count := 0
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "Pagamento Pix") {
				writer.WriteString(line + "\n")
				count++
			}
		}
		writer.Flush()
		fmt.Printf("✅ %d linhas contendo 'Pagamento Pix' foram salvas em %s\n", count, outputFilterFile)
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.Flags().StringVarP(&inputFilterFile, "input", "i", "", "Arquivo de entrada")
	filterCmd.Flags().StringVarP(&outputFilterFile, "output", "o", "", "Arquivo de saída")
}
