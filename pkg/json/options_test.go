package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		event *Event
		want  *Event
	}{
		{
			name: "add inexisting key value ok",
			args: args{
				key:   "id",
				value: "123",
			},
			event: &Event{},
			want:  &Event{"id": "123"},
		},
		{
			name: "add existing key no changes",
			args: args{
				key:   "id",
				value: "456",
			},
			event: &Event{"id": "123"},
			want:  &Event{"id": "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			got := Add(tt.args.key, tt.args.value)

			// Act
			got(tt.event)

			// Assert
			assert.Equal(t, tt.want, tt.event, "got %v, expected %v", tt.event, tt.want)
		})
	}
}

func TestReject(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		event *Event
		want  *Event
	}{
		{
			name: "reject by key value pair ok",
			args: args{
				key:   "id",
				value: "123",
			},
			event: &Event{"id": "123"},
		},
		{
			name: "reject by key value pair nothing found",
			args: args{
				key:   "id",
				value: "456",
			},
			event: &Event{"id": "123"},
			want:  &Event{"id": "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			got := Reject(tt.args.key, tt.args.value)

			// Act
			got(tt.event)

			// Assert
			if tt.want == nil {
				assert.Empty(t, tt.event)
			} else {
				assert.Equal(t, tt.want, tt.event, "got %v, expected %v", tt.event, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		event *Event
		want  *Event
	}{
		{
			name: "remove key ok",
			args: args{
				key: "id",
			},
			event: &Event{"id": "123", "title": "my title"},
			want:  &Event{"title": "my title"},
		},
		{
			name: "remove key not found ok",
			args: args{
				key: "created_at",
			},
			event: &Event{"id": "123"},
			want:  &Event{"id": "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			got := Remove(tt.args.key)

			// Act
			got(tt.event)

			// Assert
			assert.Equal(t, tt.want, tt.event, "got %v, expected %v", tt.event, tt.want)
		})
	}
}

func TestPrefix(t *testing.T) {
	type args struct {
		key    string
		prefix string
	}
	tests := []struct {
		name  string
		args  args
		event *Event
		want  *Event
	}{
		{
			name: "prefix key ok",
			args: args{
				key:    "id",
				prefix: "legacy",
			},
			event: &Event{"id": "123"},
			want:  &Event{"legacy_id": "123"},
		},
		{
			name: "prefix key not found",
			args: args{
				key:    "title",
				prefix: "legacy",
			},
			event: &Event{"id": "123"},
			want:  &Event{"id": "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			got := Prefix(tt.args.key, tt.args.prefix)

			// Act
			got(tt.event)

			// Assert
			assert.Equal(t, tt.want, tt.event, "got %v, expected %v", tt.event, tt.want)
		})
	}
}
