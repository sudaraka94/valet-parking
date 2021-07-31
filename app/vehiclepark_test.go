package app

import (
	"main/logging"
	"testing"
)

func Test_vehicleParkImpl_calculateFair(t *testing.T) {
	type fields struct {
		vehicleMap   map[string]vehicleInfo
		slotArrayMap map[string][]bool
		fairMap      map[string]float64
		logger       logging.Logger
	}
	type args struct {
		start       int
		end         int
		vehicleType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vehicleParkImpl{
				vehicleMap:   tt.fields.vehicleMap,
				slotArrayMap: tt.fields.slotArrayMap,
				fairMap:      tt.fields.fairMap,
				logger:       tt.fields.logger,
			}
			got, err := v.calculateFair(tt.args.start, tt.args.end, tt.args.vehicleType)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateFair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateFair() got = %v, want %v", got, tt.want)
			}
		})
	}
}
