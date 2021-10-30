package main

import "testing"

func TestCalculator(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "WithTwoValidArguments_Returns_Success", args: args{x: 5,y: 5,}, want: 10, wantErr: false },
		{name: "WithNegXArgument_Returns_Error", args: args{x: -1,y: 5,}, want: 0, wantErr: true },
		{name: "WithNegYArgument_Returns_Error", args: args{x: 1,y: -1,}, want: 0, wantErr: true },
		{name: "WithValidArguments_Throw_Exception", args: args{x: 5,y: 5,}, want: 5, wantErr: false },
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculator(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculator() got = %v, want %v", got, tt.want)
			}
		})
	}
}
