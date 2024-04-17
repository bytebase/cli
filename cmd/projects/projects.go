package projects

import (
	"github.com/spf13/cobra"

	"github.com/bytebase/cli/util"
)

func Cmd(s *util.Setting) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects <command>",
		Short: "List projects",
	}

	cmd.AddCommand(listCmd(s))
	return cmd
}
