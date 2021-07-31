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
			name: "CLI logger",
			args: args{
				loggerType: "cli",
			},
			want: logging.NewCLILogger(),
			wantErr: false,
		},
		{
			name: "Invalid logger",
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
			name: "Calculate normal time",
			args: args{
				start: 1613541902,
				end: 1613549740,
			},
			want: 3,
			wantErr: false,
		},
		{
			name: "Calculate with invalid inputs",
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