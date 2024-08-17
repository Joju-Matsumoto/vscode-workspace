package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Joju-Matsumoto/vscode-workspace/vscodeworkspace"
)

func ShowWorkspaces(wss ...vscodeworkspace.Workspace) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "\033[00mNAME\033[0m\tCOUNT\tPATH\n")
	for _, ws := range wss {
		fmt.Fprintf(w, "\033[36m%s\033[0m\t%5d\t%s\n", ws.Name, ws.Count, ws.Path)
	}
	w.Flush()
}

func ListWorkspaceNamesWithDescription(wss []vscodeworkspace.Workspace) []string {
	nwds := make([]string, 0, len(wss))
	for _, ws := range wss {
		nwds = append(nwds, fmt.Sprintf("%s\t%s", ws.Name, ws.Path))
	}
	return nwds
}
