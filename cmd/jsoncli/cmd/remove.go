package cmd

import (
	"errors"

	"github.com/cedric-parisi/jsoncli/pkg/json"
	"github.com/spf13/cobra"
)

// removeDmd defines the remove command
// Ensure that 1 argument is passed, no more no less
// call Process with only Remove option
var removeCmd = &cobra.Command{
	Use:     "remove <key>",
	Example: "jsoncli remove created_at",
	Short:   "remove all occurrences of the key from the input",
	Args:    cobra.RangeArgs(1, 1),
	RunE:    runRemoveCmd,
}

func runRemoveCmd(cmd *cobra.Command, args []string) (err error) {
	defer func() {
		Output.Close()
		ierr := Input.Close()
		if err == nil {
			err = ierr
		}
	}()

	if len(args) != 1 {
		return errors.New("remove requires 1 argument")
	}

	return p.Process(Input, Output, json.Remove(args[0]))
}
