package cmd

import (
	"bytes"
	"reflect"
	"testing"
)

type nopCloser struct {
	*bytes.Buffer
}

func (n *nopCloser) Close() error {
	return nil
}

func Test_validateTypeFlag(t *testing.T) {
	type args struct {
		flType string
		input  string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "validate int ok",
			args: args{
				flType: intType,
				input:  "1",
			},
			want: 1,
		},
		{
			name: "validate failed due to input not int",
			args: args{
				flType: intType,
				input:  "not_a_number",
			},
			wantErr: true,
		},
		{
			name: "validate float ok",
			args: args{
				flType: floatType,
				input:  "1.0",
			},
			want: 1.0,
		},
		{
			name: "validate failed due to input not float",
			args: args{
				flType: floatType,
				input:  "not_a_number",
			},
			wantErr: true,
		},
		{
			name: "validate string ok",
			args: args{
				flType: stringType,
				input:  "title",
			},
			want: "title",
		},
		{
			name: "validate bool ok",
			args: args{
				flType: boolType,
				input:  "true",
			},
			want: true,
		},
		{
			name: "validate failed due to input not boolean",
			args: args{
				flType: boolType,
				input:  "not_a_boolean",
			},
			wantErr: true,
		},
		{
			name: "validate failed due to unsupported type",
			args: args{
				flType: "not_a_supported_type",
				input:  "title",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateTypeFlag(tt.args.flType, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateTypeFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateTypeFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
