package cmd

import (
	"errors"

	"github.com/cedric-parisi/jsoncli/pkg/json"
	"github.com/spf13/cobra"
)

// prefixCmd defines prefix command
// ensure that 2 arguments are passed, no more no less
// call Process with only Prefix option
var prefixCmd = &cobra.Command{
	Use:     "prefix <key> <prefix>",
	Example: "jsoncli prefix id event",
	Short:   "add a prefix to a key",
	Args:    cobra.RangeArgs(2, 2),
	RunE:    runPrefixCmd,
}

func runPrefixCmd(cmd *cobra.Command, args []string) (err error) {
	defer func() {
		Output.Close()
		ierr := Input.Close()
		if err == nil {
			err = ierr
		}
	}()

	if len(args) != 2 {
		return errors.New("prefix requires 2 arguments")
	}

	return p.Process(Input, Output, json.Prefix(args[0], args[1]))
}
