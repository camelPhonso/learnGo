package db

import (
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name string
		filePath string
		want bool
	}{
		{"Valid scenario", "../test_movie_list.txt", false},
		{"Invalid scenario", "../fake_file_path.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Read(tt.filePath)

			if (err != nil) != tt.want {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}