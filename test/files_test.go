package files

import (
	"testing"

	"github.com/PonomarevAlexxander/files"
)

func TestContains(t *testing.T) {
	type args struct {
		filepath  string
		substring string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"substring found", args{"testdata/file.txt", "substring"}, true, false},
		{"substring not found", args{"testdata/file.txt", "no_such_string"}, false, false},
		{"no such file/directory test", args{"testdata/text.txt", "substring"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := files.Contains(tt.args.filepath, tt.args.substring)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
