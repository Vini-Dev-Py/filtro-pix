package main

import "github.com/Vini-Dev-Py/filtro-pix/cmd"

var version = "dev" // será substituído pelo GoReleaser com -ldflags

func main() {
	cmd.Execute(version)
}
