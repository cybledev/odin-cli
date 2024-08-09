package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// debug utility function
func validateFlags(cmd *cobra.Command) error {
	flags := cmd.Flags()

	var lastFlag string
	flags.VisitAll(func(f *pflag.Flag) {
		lastFlag = f.Name
		fmt.Println("this is the val::", f.Value.String(), "::", f.Name)
		if strings.Contains(f.Value.String(), " ") {
			parts := strings.SplitN(f.Value.String(), " ", 2)
			if len(parts) > 1 {
				cmd.SetArgs(append(cmd.Flags().Args(), parts[1]))
				f.Value.Set(parts[0])
			}
		}
	})

	if len(cmd.Flags().Args()) > 0 {
		return fmt.Errorf("unexpected arguments after '--%s' flag: %v\n"+
			"If your query contains spaces or special characters, make sure to properly quote or escape them",
			lastFlag, cmd.Flags().Args())
	}
	return nil
}
