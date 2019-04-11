package cmd

import (
	"errors"

	"github.com/cedric-parisi/jsoncli/pkg/json"
	"github.com/spf13/cobra"
)

func init() {
	// add optional flags to add command
	addCmd.Flags().StringVarP(&valueType, "type", "t", "string", "Type of value: int | float | string | bool")
}

// addCmd defines add command
// ensure that 2 arguments are passed, no more no less
// ensure that the flag type is supported otherwise return an error
// call Process with only Add option
var addCmd = &cobra.Command{
	Use:     "add <key> <value>",
	Example: "jsoncli add price 38.5 -t=float",
	Short:   "add a key value pair to the input",
	Args:    cobra.RangeArgs(2, 2),
	RunE:    runAddCmd,
}

func runAddCmd(cmd *cobra.Command, args []string) (err error) {
	defer func() {
		Output.Close()
		ierr := Input.Close()
		if err == nil {
			err = ierr
		}
	}()

	if len(args) != 2 {
		return errors.New("run requires 2 arguments")
	}

	var value interface{}
	value, err = validateTypeFlag(valueType, args[1])
	if err != nil {
		return err
	}

	// Process the add option on the input stream
	return p.Process(Input, Output, json.Add(args[0], value))
}
