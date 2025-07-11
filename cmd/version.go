package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Exibe a versão da ferramenta",
		Long:  "Mostra a versão atual embutida no binário via GoReleaser",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Versão:", rootCmd.Version)
		},
	})
}
