package cmd

import (
	"errors"

	"github.com/cedric-parisi/jsoncli/pkg/json"
	"github.com/spf13/cobra"
)

func init() {
	// add optional flags to reject command
	rejectCmd.Flags().StringVarP(&valueType, "type", "t", "string", "Type of value: int | float | string | bool")
}

// rejectCmd defines the reject command
// ensure that 2 arguments are passed, no more no less
// call Process with only Reject option
var rejectCmd = &cobra.Command{
	Use:     "reject <key> <value>",
	Example: "jsoncli reject id 4649 -t=float",
	Short:   "reject all events matching the key/value pair from the input",
	Args:    cobra.RangeArgs(2, 2),
	RunE:    runRejectCmd,
}

func runRejectCmd(cmd *cobra.Command, args []string) (err error) {
	defer func() {
		Output.Close()
		ierr := Input.Close()
		if err == nil {
			err = ierr
		}
	}()

	if len(args) != 2 {
		return errors.New("reject requires 2 arguments")
	}

	var value interface{}
	value, err = validateTypeFlag(valueType, args[1])
	if err != nil {
		return err
	}

	return p.Process(Input, Output, json.Reject(args[0], value))
}
