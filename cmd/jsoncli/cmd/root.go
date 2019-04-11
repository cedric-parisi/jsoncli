package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/cedric-parisi/jsoncli/pkg/json"
	"github.com/spf13/cobra"
)

// Input hosts the source stream
var Input io.ReadCloser

// Output hosts the destination stream
var Output io.WriteCloser

var inputFilename, outputFilename string

var p json.Processor

// Supported types for the value to add
const (
	intType       = "int"
	floatType     = "float"
	stringType    = "string"
	boolType      = "bool"
	maxByteToRead = 1024
)

var valueType string

func init() {
	// Setup processor
	p = json.NewProcessor(maxByteToRead)

	// input and output flags are persistent because they can be use from every child command
	RootCommand.PersistentFlags().StringVarP(&inputFilename, "input", "i", "", "Input stream to read from (default stdin)")
	RootCommand.PersistentFlags().StringVarP(&outputFilename, "output", "o", "", "Output to write to (default stdout)")

	// Add child commands to jsoncli
	RootCommand.AddCommand(addCmd)
	RootCommand.AddCommand(prefixCmd)
	RootCommand.AddCommand(rejectCmd)
	RootCommand.AddCommand(removeCmd)
}

// RootCommand defines the root jsoncli command
var RootCommand = &cobra.Command{
	Use:   "jsoncli",
	Short: `Executes command against a JSON stream input and output it`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Set to STDIN and STDOUT by default
		Input = os.Stdin
		Output = os.Stdout

		// PeristentPreRunE is executed before the RunE method
		// children inherit and execute this method
		// If the flag input is provided
		if inputFilename != "" {
			var err error
			// Open the source file
			Input, err = os.Open(inputFilename)
			if err != nil {
				return err
			}
		}

		// If the flag output is provided
		if outputFilename != "" {
			var err error
			// Create the destination file
			Output, err = os.Create(outputFilename)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

// ensure that the type flag is supported
// and return the type flag value as interface{}
func validateTypeFlag(flType, input string) (interface{}, error) {
	if flType != stringType &&
		flType != intType &&
		flType != floatType &&
		flType != boolType {
		return nil, fmt.Errorf("type %s not supported", flType)
	}

	// Convert the value into the requested type
	// otherwise cli args are all treated as string
	var value interface{}
	var err error
	switch flType {
	case "int":
		value, err = strconv.Atoi(input)
		if err != nil {
			return nil, fmt.Errorf("could not parse %s to int", input)
		}
		break
	case "float":
		value, err = strconv.ParseFloat(input, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse %s to float", input)
		}
		break
	case "bool":
		value, err = strconv.ParseBool(input)
		if err != nil {
			return nil, fmt.Errorf("could not parse %s to bool", input)
		}
		break
	case "string":
		value = input
		break
	}
	return value, nil
}
