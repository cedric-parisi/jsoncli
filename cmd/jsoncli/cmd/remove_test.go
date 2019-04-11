package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_runRemoveCmd(t *testing.T) {
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
			name: "run remove cmd ok",
			args: args{
				args: []string{"id"},
			},
			input: `{"id":"123","title":"my title"}`,
			want:  `{"title":"my title"}`,
		},
		{
			name: "run remove failed due to too many arguments",
			args: args{
				args: []string{"title", "id"},
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

			if err := runRemoveCmd(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("runRemoveCmd() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Assert
			if !tt.wantErr {
				assert.Equal(t, tt.want, o.String())
			}
		})
	}
}
