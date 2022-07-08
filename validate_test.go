package ossinspector

import "testing"

func Test_checkExpr(t *testing.T) {
	type args struct {
		checkString string
		value       uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "true test", args: args{checkString: ">100", value: 101}, want: true},
		{name: "false test", args: args{checkString: "<100", value: 90}, want: true},
		{name: "false test2", args: args{checkString: "<100", value: 1000}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkExpr(tt.args.checkString, tt.args.value); got != tt.want {
				t.Errorf("checkExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}