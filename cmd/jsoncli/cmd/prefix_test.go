package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_runPrefixCmd(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		args    args
		input   string
		wantErr bool
		want    string
	}{
		{
			name: "run prefix cmd ok",
			args: args{
				args: []string{"id", "legacy"},
			},
			input: `{"id":"test"}`,
			want:  `{"legacy_id":"test"}`,
		},
		{
			name: "prefix add failed due not enough arguments",
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

			// Act
			if err := runPrefixCmd(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("runPrefixCmd() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Assert
			if !tt.wantErr {
				assert.Equal(t, tt.want, o.String())
			}
		})
	}
}
