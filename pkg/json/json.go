package json

import (
	"encoding/json"
	"fmt"
	"io"
)

// Event defines a json object from the input source
// We use map instead of struct to be able to process any kind of JSON object
// and to be able to perform actions on the fields (add, remove, ...)
type Event map[string]interface{}

type processor struct {
	limit int64
}

// Processor ...
type Processor interface {
	Process(input io.Reader, output io.Writer, options ...Option) error
}

// NewProcessor creates a new processor
// limit is the max number of bytes allowed to read
// set to 0 to have unlimited reading
func NewProcessor(limit int64) Processor {
	return &processor{
		limit: limit,
	}
}

// Process runs all options against the input and writes to the output
func (p *processor) Process(input io.Reader, output io.Writer, options ...Option) error {
	limitedInput := input

	// If a limit was set at construction
	if p.limit > 0 {
		limitedInput = io.LimitReader(input, p.limit)
	}
	dec := json.NewDecoder(limitedInput)

	// We loop until there is no more JSON object to decode
	for dec.More() {
		event := &Event{}
		if err := dec.Decode(event); err != nil {
			// LimitReader returns an error when limit reached in the middle of a block
			// LimitReader returns EOF when n bytes read
			if err == io.ErrUnexpectedEOF || err == io.EOF {
				return fmt.Errorf("limit of %d reached", p.limit)
			}
			return err
		}

		// Execute every options
		for _, opt := range options {
			opt(event)
		}

		if event != nil && len(*event) > 0 {
			// Encode the event to JSON
			line, err := json.Marshal(event)
			if err != nil {
				return err
			}

			// Write the JSON to the output
			if _, err := output.Write(line); err != nil {
				return err
			}
		}
	}
	return nil
}
