package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filtro-pix",
	Short: "Ferramenta CLI para processar logs de pagamento Pix",
	Long:  `Uma CLI simples que permite filtrar logs Pix, extrair TXIDs e visualizar a versão.`,
}

// Execute inicia o comando raiz com suporte à versão
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
