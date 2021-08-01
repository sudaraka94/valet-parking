package app

import (
	"main/logging"
	"main/util"
	"reflect"
	"testing"
)

func Test_vehicleParkImpl_calculateFare(t *testing.T) {
	type fields struct {
		vehicleMap   map[string]vehicleInfo
		slotArrayMap map[string][]bool
		fareMap      map[string]float64
		logger       logging.Logger
	}
	type args struct {
		start       int
		end         int
		vehicleType string
	}
	tests := []struct {
		name    					string
		fields  					fields
		args    					args
		want    					float64
		wantErr 					bool
		calculateTimeMethodMock		func(int, int) (int, error)
	}{
		{
			name: "Should calculate fare properly",
			fields: fields{
				fareMap: map[string]float64{
					"car": float64(1),
				},
			},
			args: args{
				vehicleType: "car",
			},
			want: 2,
			wantErr: false,
			calculateTimeMethodMock: func(start int, end int) (int, error) {
				return 2, nil
			},
		},
		{
			name: "Should throw invalid vehicle type error",
			fields: fields{
				fareMap: map[string]float64{
					"car": float64(1),
				},
			},
			args: args{
				vehicleType: "carx",
			},
			want: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vehicleParkImpl{
				vehicleMap:   tt.fields.vehicleMap,
				slotArrayMap: tt.fields.slotArrayMap,
				fareMap:      tt.fields.fareMap,
				logger:       tt.fields.logger,
			}
			// Mocking the method
			if tt.calculateTimeMethodMock != nil {
				calculateTime = tt.calculateTimeMethodMock
			} else {
				calculateTime = util.CalculateTime
			}

			got, err := v.calculateFare(tt.args.start, tt.args.end, tt.args.vehicleType)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateFare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateFare() got = %v, want %v", got, tt.want)
			}
		})
	}
	// reset the mocked method at the end
	calculateTime = util.CalculateTime
}

func Test_vehicleParkImpl_addVehicle(t *testing.T) {
	type fields struct {
		vehicleMap   map[string]vehicleInfo
		slotArrayMap map[string][]bool
		fareMap      map[string]float64
		logger       logging.Logger
	}
	type args struct {
		vehicleType string
		regNo       string
		timestamp   int
	}
	tests := []struct {
		name    					string
		fields  					fields
		args    					args
		want    					int
		wantErr 					bool
		wantFields					fields
		mockedFindAndUpdateMethod	func([]bool) int
	}{
		{
			name: "Should insert a vehicle",
			fields: fields{
				vehicleMap: map[string]vehicleInfo{},
				slotArrayMap: map[string][]bool{
					"car": {},
				},
			},
			args: args{
				vehicleType: "car",
				regNo: "x",
				timestamp: 2,
			},
			want: 0,
			wantErr: false,
			wantFields: fields{
				vehicleMap: map[string]vehicleInfo{"x": {
					SlotNumber:  0,
					Timestamp:   2,
					VehicleType: "car",
				}},
			},
			mockedFindAndUpdateMethod: func(arr []bool) int {
				return 0
			},
		},
		{
			name: "Should reject the vehicle",
			fields: fields{
				vehicleMap: map[string]vehicleInfo{},
				slotArrayMap: map[string][]bool{
					"car": []bool{},
				},
			},
			args: args{
				vehicleType: "car",
				regNo: "x",
				timestamp: 2,
			},
			want: -1,
			wantErr: false,
			wantFields: fields{
				vehicleMap: map[string]vehicleInfo{},
			},
			mockedFindAndUpdateMethod: func(arr []bool) int {
				return -1
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vehicleParkImpl{
				vehicleMap:   tt.fields.vehicleMap,
				slotArrayMap: tt.fields.slotArrayMap,
				fareMap:      tt.fields.fareMap,
				logger:       tt.fields.logger,
			}

			if tt.mockedFindAndUpdateMethod != nil {
				findAndUpdate = tt.mockedFindAndUpdateMethod
			} else {
				findAndUpdate = util.FindAndUpdate
			}

			got, err := v.addVehicle(tt.args.vehicleType, tt.args.regNo, tt.args.timestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("addVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("addVehicle() got = %v, want %v", got, tt.want)
			}
			gotFields := fields{
				vehicleMap: v.vehicleMap,
			}
			if !reflect.DeepEqual(fields{}, tt.wantFields) && !reflect.DeepEqual(gotFields, tt.wantFields) {
				t.Errorf("addVehicle() got = %v, want %v", gotFields, tt.wantFields)
			}
		})
		// cleanup mocks
		findAndUpdate = util.FindAndUpdate
	}
}

func Test_vehicleParkImpl_removeVehicle(t *testing.T) {
	type fields struct {
		vehicleMap   map[string]vehicleInfo
		slotArrayMap map[string][]bool
		fareMap      map[string]float64
		logger       logging.Logger
	}
	type args struct {
		regNo string
	}
	tests := []struct {
		name   			string
		fields 			fields
		args   			args
		want   			vehicleInfo
		wantFields		fields
		wantErr			bool
	}{
		{
			name: "Should remove the car",
			fields: fields{
				vehicleMap: map[string]vehicleInfo{"x": {SlotNumber: 0, VehicleType: "car"}},
				slotArrayMap: map[string][]bool{
					"car": {true},
				},
			},
			args: args{
				regNo: "x",
			},
			want: vehicleInfo{SlotNumber: 0, VehicleType: "car"},
			wantFields: fields{
				vehicleMap: map[string]vehicleInfo{},
				slotArrayMap: map[string][]bool{
					"car": []bool{false},
				},
			},
			wantErr: false,
		},
		{
			name: "Should throw error",
			fields: fields{
				vehicleMap: map[string]vehicleInfo{},
			},
			args: args{
				regNo: "x",
			},
			wantErr: true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vehicleParkImpl{
				vehicleMap:   tt.fields.vehicleMap,
				slotArrayMap: tt.fields.slotArrayMap,
				fareMap:      tt.fields.fareMap,
				logger:       tt.fields.logger,
			}
			got, err := v.removeVehicle(tt.args.regNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("addVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeVehicle() = %v, want %v", got, tt.want)
			}
			gotFields := fields{
				slotArrayMap: v.slotArrayMap,
				vehicleMap: v.vehicleMap,
			}
			if !reflect.DeepEqual(fields{}, tt.wantFields) && !reflect.DeepEqual(gotFields, tt.wantFields) {
				t.Errorf("addVehicle() got = %v, want %v", gotFields, tt.wantFields)
			}
		})
	}
}