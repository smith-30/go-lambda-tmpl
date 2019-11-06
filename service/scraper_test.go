package service

import (
	"testing"
	"time"
)

func TestListScrape(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListScrape()
		})
	}
}

func TestThisFriday(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				t: time.Date(2019, 11, 05, 12, 00, 00, 0, time.Local),
			},
			want: "20191108",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThisFriday(tt.args.t, "20191108"); got != tt.want {
				t.Errorf("ThisFriday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWheather(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wheather(); got != tt.want {
				t.Errorf("Wheather() = %v, want %v", got, tt.want)
			}
		})
	}
}
