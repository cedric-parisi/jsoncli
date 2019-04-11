package json

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		input   io.Reader
		options []Option
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "execute ok with no options",
			args: args{
				input: strings.NewReader(`{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"lng":2.0403327446430257}`),
			},
			wantOutput: `{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"lng":2.0403327446430257}`,
		},
		{
			name: "execute ok with add option",
			args: args{
				input: strings.NewReader(`{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"lng":2.0403327446430257}`),
				options: []Option{
					Add("updated_at", "2019-03-27 19:00:00"),
				},
			},
			wantOutput: `{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"lng":2.0403327446430257,"updated_at":"2019-03-27 19:00:00"}`,
		},
		{
			name: "execute ok with prefix option",
			args: args{
				input: strings.NewReader(`{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"lng":2.0403327446430257}`),
				options: []Option{
					Prefix("lng", "pos"),
				},
			},
			wantOutput: `{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"pos_lng":2.0403327446430257}`,
		},
		{
			name: "execute ok with remove option",
			args: args{
				input: strings.NewReader(`{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"lng":2.0403327446430257}`),
				options: []Option{
					Remove("id"),
				},
			},
			wantOutput: `{"created_at":"2016-12-14 07:00:00","lat":49.01249051526539,"lng":2.0403327446430257}`,
		},
		{
			name: "execute ok with reject option",
			args: args{
				input: strings.NewReader(`{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"lng":2.0403327446430257}`),
				options: []Option{
					Reject("lat", 49.01249051526539),
				},
			},
			wantOutput: "",
		},
		{
			name: "execute ok with all options",
			args: args{
				input: strings.NewReader(`
					{"created_at":"2016-12-14 07:00:00","id":4649,"lat":49.01249051526539,"lng":2.0403327446430257}
					{"created_at":"2016-12-14 07:00:00","id":10086,"lat":48.907344373066344,"lng":2.3638633128958166}`),
				options: []Option{
					Add("title", "my title"),
					Remove("id"),
					Reject("lat", 49.01249051526539),
					Prefix("lng", "pos"),
				},
			},
			wantOutput: `{"created_at":"2016-12-14 07:00:00","lat":48.907344373066344,"pos_lng":2.3638633128958166,"title":"my title"}`,
		},
		{
			name: "execute ko due to invalid json",
			args: args{
				input: strings.NewReader(`{"not_a_valid_json"}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			if err := Process(tt.args.input, output, tt.args.options...); (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("Process() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
