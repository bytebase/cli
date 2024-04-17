package projects

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/bytebase/cli/util"

	v1pb "buf.build/gen/go/bytebase/bytebase/protocolbuffers/go/v1"
)

// ListCmd encapsulates the command for listing backups for a branch.
func listCmd(s *util.Setting) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all projects",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			slog.Info("listing all projects...\n")
			resp, err := s.Config.ProjectServiceClient.ListProjects(ctx, &v1pb.ListProjectsRequest{})
			if err != nil {
				return err
			}
			projects := resp.Projects

			for _, project := range projects {
				slog.Info(fmt.Sprintf("project %s", project.Name)) // Added format specifier for project name
			}
			return nil
		},
	}

	return cmd
}
