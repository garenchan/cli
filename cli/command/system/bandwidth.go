package system

import (
	"context"
	"fmt"

	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

// BandwidthSetOptions specifies some options that are used when setting bandwidth.
type BandwidthSetOptions struct {
	Value      int64
	Persistent bool
}

// newBandwidthCommand creates a new cobra.Command for `docker system bandwidth`
func newBandwidthCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bandwidth <subcommand>",
		Short: "bandwidth related commands",
	}

	cmd.AddCommand(newDownloadBandwidthCommand(dockerCli))
	cmd.AddCommand(newUploadBandwidthCommand(dockerCli))

	return cmd
}

// newBandwidthGetCommand returns the cobra command for "bandwidth download".
func newDownloadBandwidthCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download <subcommand>",
		Short: "download bandwidth related commands",
	}

	cmd.AddCommand(newDownloadBandwidthGetCommand(dockerCli))
	cmd.AddCommand(newDownloadBandwidthSetCommand(dockerCli))

	return cmd
}

// newDownloadBandwidthGetCommand returns the cobra command for "bandwidth download get".
func newDownloadBandwidthGetCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get Max Download Bandwidth",

		RunE: func(cmd *cobra.Command, args []string) error {
			return runDownloadBandwidthGet(cmd, dockerCli)
		},
	}

	return cmd
}

// newDownloadBandwidthSetCommand returns the cobra command for "bandwidth download set".
func newDownloadBandwidthSetCommand(dockerCli command.Cli) *cobra.Command {
	bandwidth := &BandwidthSetOptions{Value: -1, Persistent: false}

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set Max Download Bandwidth",

		RunE: func(cmd *cobra.Command, args []string) error {
			return runDownloadBandwidthSet(cmd, dockerCli, bandwidth)
		},
	}

	flags := cmd.Flags()
	flags.Int64VarP(&bandwidth.Value, "value", "v", bandwidth.Value, "Bandwidth value")
	flags.BoolVarP(&bandwidth.Persistent, "persistent", "p", bandwidth.Persistent, "Whether to persist")

	return cmd
}

// newUploadBandwidthCommand returns the cobra command for "bandwidth upload".
func newUploadBandwidthCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload <subcommand>",
		Short: "upload bandwidth related commands",
	}

	cmd.AddCommand(newUploadBandwidthGetCommand(dockerCli))
	cmd.AddCommand(newUploadBandwidthSetCommand(dockerCli))

	return cmd
}

// newUploadBandwidthGetCommand returns the cobra command for "bandwidth upload get".
func newUploadBandwidthGetCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get Max Upload Bandwidth",

		RunE: func(cmd *cobra.Command, args []string) error {
			return runUploadBandwidthGet(cmd, dockerCli)
		},
	}

	return cmd
}

// newUploadBandwidthSetCommand returns the cobra command for "bandwidth upload set".
func newUploadBandwidthSetCommand(dockerCli command.Cli) *cobra.Command {
	bandwidth := &BandwidthSetOptions{Value: -1, Persistent: false}

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set Max Upload Bandwidth",

		RunE: func(cmd *cobra.Command, args []string) error {
			return runUploadBandwidthSet(cmd, dockerCli, bandwidth)
		},
	}

	flags := cmd.Flags()
	flags.Int64VarP(&bandwidth.Value, "value", "v", bandwidth.Value, "Bandwidth value")
	flags.BoolVarP(&bandwidth.Persistent, "persistent", "p", bandwidth.Persistent, "Whether to persist")

	return cmd
}

func runDownloadBandwidthGet(cmd *cobra.Command, dockerCli command.Cli) error {
	ctx := context.Background()
	if bandwidth, err := dockerCli.Client().GetDownloadBandwidth(ctx); err != nil {
		return err
	} else {
		fmt.Println(bandwidth)
	}

	return nil
}

func runDownloadBandwidthSet(cmd *cobra.Command, dockerCli command.Cli, options *BandwidthSetOptions) error {
	ctx := context.Background()
	if err := dockerCli.Client().SetDownloadBandwidth(ctx, options.Value, options.Persistent); err != nil {
		return err
	}

	return nil
}

func runUploadBandwidthGet(cmd *cobra.Command, dockerCli command.Cli) error {
	ctx := context.Background()
	if bandwidth, err := dockerCli.Client().GetUploadBandwidth(ctx); err != nil {
		return err
	} else {
		fmt.Println(bandwidth)
	}

	return nil
}

func runUploadBandwidthSet(cmd *cobra.Command, dockerCli command.Cli, options *BandwidthSetOptions) error {
	ctx := context.Background()
	if err := dockerCli.Client().SetUploadBandwidth(ctx, options.Value, options.Persistent); err != nil {
		return err
	}

	return nil
}
