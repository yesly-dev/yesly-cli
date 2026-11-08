package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCommand writes the Markdown reference in docs/commands. Hidden: it is
// for this repository, not for users.
func docsCommand(root *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:    "gen-docs <dir>",
		Short:  "Write the Markdown command reference",
		Hidden: true,
		Args:   cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := args[0]
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return err
			}
			old, _ := filepath.Glob(filepath.Join(dir, "yesly*.md"))
			for _, f := range old {
				if err := os.Remove(f); err != nil {
					return err
				}
			}
			root.DisableAutoGenTag = true
			link := func(name string) string { return strings.ToLower(name) }
			if err := doc.GenMarkdownTreeCustom(root, dir, func(string) string { return "" }, link); err != nil {
				return err
			}
			fmt.Fprintf(cmd.ErrOrStderr(), "wrote %s\n", dir)
			return nil
		},
	}
}
