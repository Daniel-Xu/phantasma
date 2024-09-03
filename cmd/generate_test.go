package cmd

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "Valid date",
			input:   "2023-04-15",
			want:    time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "Invalid date format",
			input:   "15-04-2023",
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "Invalid date",
			input:   "2023-13-32",
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "Empty string",
			input:   "",
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("parseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
