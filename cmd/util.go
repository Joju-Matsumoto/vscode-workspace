package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Joju-Matsumoto/vscode-workspace/vscodeworkspace"
)

func ShowWorkspaces(wss ...vscodeworkspace.Workspace) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, ws := range wss {
		fmt.Fprintf(w, "\033[36m%s\033[0m\t%s\n", ws.Name, ws.Path)
	}
	w.Flush()
}
