package benchmark

import (
	"testing"
)

func TestNewCaller(t *testing.T) {
	type args struct {
		fn callerFunc
	}
	tests := []struct {
		name string
		args args
		want *caller
	}{
		// TODO: Add test cases.
		{"noop", args{fn: CallSetGlobalReturnValue}, &caller{fn: CallSetGlobalReturnValue, fnTrue: CallSetGlobalReturnValue, fnFalse: noop}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &caller{fn: tt.args.fn, fnTrue: tt.args.fn, fnFalse: noop}

			gFn := NewAnyValue(c.fn)
			wFn := NewAnyValue(tt.want.fn)

			AssertSameType(t, tt.name, gFn, wFn)
			AssertSameKind(t, tt.name, gFn, wFn)

			CompareFuncs(t, tt.name, gFn, wFn)

		})
	}
}
