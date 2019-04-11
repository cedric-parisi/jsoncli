package main

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "add ok",
			args:  []string{"add", "title", "my title"},
			input: `{"id":"123"}`,
			want:  `{"id":"123","title":"my title"}`,
		},
		{
			name:    "add failed due to not enough args",
			args:    []string{"add", "title"},
			input:   `{"id":"123"}`,
			wantErr: true,
		},
		{
			name:  "reject ok",
			args:  []string{"reject", "id", "123"},
			input: `{"id":"123"}`,
			want:  ``,
		},
		{
			name:    "reject failed due to not enough args",
			args:    []string{"reject", "id"},
			input:   `{"id":"123"}`,
			wantErr: true,
		},
		{
			name:  "remove ok",
			args:  []string{"remove", "id"},
			input: `{"id":"123"}`,
			want:  ``,
		},
		{
			name:    "remove failed due to too many args",
			args:    []string{"remove", "id", "title"},
			input:   `{"id":"123"}`,
			wantErr: true,
		},
		{
			name:  "prefix ok",
			args:  []string{"prefix", "id", "legacy"},
			input: `{"id":"123"}`,
			want:  `{"legacy_id":"123"}`,
		},
		{
			name:    "prefix failed due to not enough args",
			args:    []string{"prefix", "id"},
			input:   `{"id":"123"}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			main()
			cmd := exec.Command("jsoncli", tt.args...)
			cmd.Stdin = strings.NewReader(tt.input)

			// Act
			res, err := cmd.Output()

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("error %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, string(res))
			}
		})
	}
}
