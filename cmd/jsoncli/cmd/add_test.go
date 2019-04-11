package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spf13/cobra"
)

func Test_runAddCmd(t *testing.T) {
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
			name: "run add cmd ok",
			args: args{
				args:      []string{"title", "my title"},
				valueType: "string",
			},
			input: `{"id":"123"}`,
			want:  `{"id":"123","title":"my title"}`,
		},
		{
			name: "run add failed due to mismatch valueType",
			args: args{
				args:      []string{"title", "my title"},
				valueType: "int",
			},
			input:   `{"id":"123"}`,
			wantErr: true,
		},
		{
			name: "run add failed due not enough arguments",
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

			// Act
			if err := runAddCmd(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("runAddCmd() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Assert
			if !tt.wantErr {
				assert.Equal(t, tt.want, o.String())
			}
		})
	}
}
