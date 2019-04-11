package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_runRejectCmd(t *testing.T) {
	type args struct {
		cmd       *cobra.Command
		valueType string
		args      []string
	}
	tests := []struct {
		name    string
		args    args
		input   string
		wantErr bool
		want    string
	}{
		{
			name: "run reject cmd ok",
			args: args{
				args:      []string{"id", "123"},
				valueType: "string",
			},
			input: `{"id":"123"}{"id":"456"}`,
			want:  `{"id":"456"}`,
		},
		{
			name: "run reject failed due to mismatch valueType",
			args: args{
				args:      []string{"id", "abc"},
				valueType: "int",
			},
			input:   `{"id":"123"}`,
			wantErr: true,
		},
		{
			name: "reject failed due not enough arguments",
			args: args{
				args: []string{"title"},
			},
			input:   `{"id":"123"}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			Input = ioutil.NopCloser(strings.NewReader(tt.input))
			o := &bytes.Buffer{}
			Output = &nopCloser{o}
			valueType = tt.args.valueType

			if err := runRejectCmd(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("runRejectCmd() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Assert
			if !tt.wantErr {
				assert.Equal(t, tt.want, o.String())
			}
		})
	}
}
