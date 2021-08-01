package app

import (
	config "main/config"
	"main/logging"
	"reflect"
	"testing"
)

func Test_appImpl_parseDataContent(t *testing.T) {
	type fields struct {
		config     *config.Config
		logger     logging.Logger
		operations []string
		slotArrays map[string][]bool
	}
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string][]bool
		want1   []string
		wantErr bool
	}{
		{
			name: "Should parse data properly",
			fields: fields{
				config: &config.Config{
					VehicleTypes: []config.VehicleType{
						{
							Name: "car",
						},
						{
							Name: "motorcycle",
						},
					},
				},
			},
			args: args{
				content: "2 2\nfirst operation",
			},
			want: map[string][]bool{
				"car": {false, false},
				"motorcycle": {false, false},
			},
			want1: []string{"first operation"},
			wantErr: false,
		},
		{
			name: "Should throw invalid data format error",
			fields: fields{
				config: &config.Config{
					VehicleTypes: []config.VehicleType{
						{
							Name: "car",
						},
						{
							Name: "motorcycle",
						},
					},
				},
			},
			args: args{
				content: "2\nfirst operation",
			},
			want: map[string][]bool{},
			want1: []string{},
			wantErr: true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &appImpl{
				config:     tt.fields.config,
				logger:     tt.fields.logger,
				operations: tt.fields.operations,
			}
			got, got1, err := a.parseDataContent(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDataContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDataContent() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseDataContent() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
