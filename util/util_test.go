package util

import (
	"main/logging"
	"reflect"
	"testing"
)

func Test_initializeLogger(t *testing.T) {
	type args struct {
		loggerType string
	}
	tests := []struct {
		name    string
		args    args
		want    logging.Logger
		wantErr bool
	}{
		{
			name: "Should return a CLI Logger",
			args: args{
				loggerType: "cli",
			},
			want: logging.NewCLILogger(),
			wantErr: false,
		},
		{
			name: "Should throw invalid logger type error",
			args: args{
				loggerType: "clix",
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitializeLogger(tt.args.loggerType)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitializeLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeLogger() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateTime(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Should calculate time properly",
			args: args{
				start: 1613541902,
				end: 1613549740,
			},
			want: 3,
			wantErr: false,
		},
		{
			name: "Should throw error on invalid timestamps",
			args: args{
				start: 1613549740,
				end: 1613541902,
			},
			want: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTime(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAndUpdate(t *testing.T) {
	type args struct {
		arr []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should process empty array properly",
			args: args{
				arr: []bool{},
			},
			want: -1,
		},
		{
			name: "Should process an array with empty slot properly",
			args: args{
				arr: []bool{false, false},
			},
			want: 0,
		},
		{
			name: "Should process an array without empty slots",
			args: args{
				arr: []bool{true, true},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAndUpdate(tt.args.arr); got != tt.want {
				t.Errorf("FindAndUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}