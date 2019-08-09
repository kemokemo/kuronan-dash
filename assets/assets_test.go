package assets

import "testing"

func TestLoadAssets(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"LoadAssets test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadAssets(); (err != nil) != tt.wantErr {
				t.Errorf("LoadAssets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCloseAssets(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"CloseAssets test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CloseAssets(); (err != nil) != tt.wantErr {
				t.Errorf("CloseAssets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
