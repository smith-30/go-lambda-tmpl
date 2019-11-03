package service

import (
	"fmt"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	tests := []struct {
		name    string
		want    *LINEToken
		wantErr bool
	}{
		{
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%#v\n", got)
		})
	}
}
