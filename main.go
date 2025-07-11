package main

import (
	"filtro-pix/cmd"
)

var version = "dev" // será substituído pelo GoReleaser com -ldflags

func main() {
	cmd.Execute(version)
}
