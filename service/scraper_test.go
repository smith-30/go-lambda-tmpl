package service

import "testing"

func TestExampleScrape(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip()
			ExampleScrape()
		})
	}
}
