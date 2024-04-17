// Package cmd is the command surface of Bytebase bb tool provided by bytebase.com.
package cmd

import (
	"context"

	"log/slog"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/bytebase/cli/cmd/projects"
	"github.com/bytebase/cli/cmd/version"
	"github.com/bytebase/cli/config"
	"github.com/bytebase/cli/util"

	rpc "buf.build/gen/go/bytebase/bytebase/grpc/go/v1/bytebasev1grpc"
)

var (
	flags struct {
		url   string
		token string
	}

	grpcConn *grpc.ClientConn

	rootCmd = &cobra.Command{
		Use:   "bb",
		Short: "CLI for Bytebase",
		Long:  "bb is the CLI interacting with https://api.bytebase.com.",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			start()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			close()
		},
	}

	s = &util.Setting{
		Config: config.New(),
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&flags.url,
		"url", "", "The URL for the Bytebase instance.")
	rootCmd.PersistentFlags().StringVar(&flags.token,
		"token", "", "The API token to interact with the Bytebase API.")
}

func start() {
	var err error
	grpcConn, err = grpc.Dial(flags.url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(unaryInterceptor(flags.token)))
	if err != nil {
		slog.Error("failed to dial grpc: %v", err)
		return
	}

	s.Config.ProjectServiceClient = rpc.NewProjectServiceClient(grpcConn)
}

func close() {
	if grpcConn != nil {
		if err := grpcConn.Close(); err != nil {
			slog.Error("failed to close grpc connection: %v", err)
		} else {
			slog.Info("grpc connection closed successfully")
		}
	}
}

// unaryInterceptor injects the access token into the request metadata
func unaryInterceptor(accessToken string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Create a new context with the access token attached
		newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+accessToken)

		// Forward the call to the invoker
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func runCmd(ctx context.Context) error {
	rootCmd.AddCommand(version.Cmd(), projects.Cmd(s))

	return rootCmd.ExecuteContext(ctx)
}

// Execute is the execute command for root command.
func Execute(ctx context.Context) int {
	err := runCmd(ctx)
	if err == nil {
		return 0
	}
	return 1
}
