package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/danehans/gonx/client"
)

var (
	// RootCmd is the base nxctl command.
	RootCmd = &cobra.Command{
		Use:   "nxctl",
		Short: "A command line client for nexus switches.",
		Long: `A CLI for Nexus switches

To get help about a resource or command, run "nxctl help resource"`,
	}

	// globalFlags can be set for any subcommand.
	globalFlags = struct {
		endpoints []string
		username  string
		password  string
	}{}
)

func init() {
	RootCmd.PersistentFlags().StringSliceVar(&globalFlags.endpoints, "endpoints", []string{"10.10.10.254"}, "IP's of Nexus switches")
	// Username and password for authenticating to the Nexus switches
	RootCmd.PersistentFlags().StringVar(&globalFlags.username, "username", "admin", "Username for authentication")
	RootCmd.PersistentFlags().StringVar(&globalFlags.password, "password", "cisco", "Password for authentication")
	cobra.EnablePrefixMatching = true
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// mustClientFromCmd returns an gonx client or exits.
func mustClientFromCmd(cmd *cobra.Command) *client.Client {
	endpoints := endpointsFromCmd(cmd)
	username := usernameFromCmd(cmd)
	password := passwordFromCmd(cmd)

	cfg := &client.Config{
		Endpoints: endpoints,
		Username:  username,
		Password:  password,
	}

	// gonx client
	client, err := client.New(cfg)
	if err != nil {
		exitWithError(ExitBadConnection, err)
	}
	return client
}

// endpointsFromCmd returns the endpoint arguments.
func endpointsFromCmd(cmd *cobra.Command) []string {
	endpoints, err := cmd.Flags().GetStringSlice("endpoints")
	if err != nil {
		exitWithError(ExitBadArgs, err)
	}
	return endpoints
}

// usernameFromCmd returns the endpoint arguments.
func usernameFromCmd(cmd *cobra.Command) string {
	username, err := cmd.Flags().GetString("username")
	if err != nil {
		exitWithError(ExitBadArgs, err)
	}
	return username
}

// passwordFromCmd returns the password arguments.
func passwordFromCmd(cmd *cobra.Command) string {
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		exitWithError(ExitBadArgs, err)
	}
	return password
}
