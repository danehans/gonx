package cli

import (
	"io/ioutil"

	"github.com/spf13/cobra"
)

// profilePutCmd creates and updates Profiles.
var (
	configPutCmd = &cobra.Command{
		Use:   "create --file FILENAME",
		Short: "Create a Nexus switch configuration",
		Long:  `Create a Nexus switch configuration`,
		Run:   runConfigPutCmd,
	}
	flagFilename string
)

func init() {
	configCmd.AddCommand(profilePutCmd)
	configPutCmd.Flags().StringVarP(&flagFilename, "filename", "f", "", "filename to use to configure a switch")
	configPutCmd.MarkFlagRequired("filename")
	configPutCmd.MarkFlagFilename("filename", "json")
}

func runConfigPutCmd(cmd *cobra.Command, args []string) {
	if len(flagFilename) == 0 {
		cmd.Help()
		return
	}
	if err := validateArgs(cmd, args); err != nil {
		return
	}

	client := mustClientFromCmd(cmd)
	config, err := loadConfig(flagFilename)
	if err != nil {
		exitWithError(ExitError, err)
	}
	req := &pb.ProfilePutRequest{Profile: profile}
	_, err = client.Profiles.ProfilePut(context.TODO(), req)
	if err != nil {
		exitWithError(ExitError, err)
	}
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return usageError(cmd, "Unexpected args: %v", args)
	}
	return nil
}

func loadConfig(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
