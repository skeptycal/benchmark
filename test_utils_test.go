package benchmark

import (
	"reflect"
	"testing"
)

func Test_limitTestResultLength(t *testing.T) {
	type args struct {
		v Any
	}
	tests := []struct {
		name   string
		in     string
		enable bool
		want   string
	}{
		{"short", "short", true, "short"},
		{"long(off)", "longlonglonglonglonglong", false, "longlonglonglonglonglong"},
		{"long(on)", "longlonglonglonglonglong", true, "longlonglong..."},
	}
	for _, tt := range tests {
		LimitResult = tt.enable
		tRun(t, tt.name, limitTestResultLength(tt.in), tt.want)
	}
}

func Test_typeGuardExclude(t *testing.T) {
	type args struct {
		needle     Any
		notAllowed []Any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"noPtr", args{reflect.Int, []Any{reflect.Ptr}}, true},
		{"noPtr", args{reflect.Int, []Any{reflect.Int}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typeGuardExclude(tt.args.needle, tt.args.notAllowed); got != tt.want {
				t.Errorf("typeGuardExclude() = %v, want %v", got, tt.want)
			}
		})
	}
}
