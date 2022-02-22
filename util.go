package benchmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/skeptycal/types"
)

var (
	LimitResult            bool
	DefaultTestResultLimit = 15

	Contains = types.Contains
)

func chr(c byte) string {
	return fmt.Sprintf("%c", c)
}

// func limitTestResultLength(v Any) string {
// 	s := fmt.Sprintf("%v", v)

// 	if len(s) > DefaultTestResultLimit && LimitResult {
// 		return s[:DefaultTestResultLimit-3] + "..."
// 	}

// 	return s
// }

// func tName(testname, itemname string) string {
// 	return fmt.Sprintf("%v.%v()", testname, itemname)
// }

// func typeGuardExclude(needle Any, notAllowed []Any) bool {
// 	return !Contains(needle, notAllowed)
// }

func tErrorf(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v(%T), want %v(%T)", name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func tRunTest(t *testing.T, tt *test) {
	if NewAnyValue(tt.got).IsComparable() && NewAnyValue(tt.want).IsComparable() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want != tt.wantErr {
				if reflect.DeepEqual(tt.got, tt.want) == tt.wantErr {
					tError(t, tt.name, tt.got, tt.want)
				}
			}
		})
	}
}

func limitTestResultLength(v Any) string {
	s := fmt.Sprintf("%v", v)

	if len(s) > DefaultTestResultLimit && LimitResult {
		return s[:DefaultTestResultLimit-3] + "..."
	}

	return s
}

func tName(testname, funcname, argname Any) string {
	if argname == "" {
		return fmt.Sprintf("%v: %v()", testname, funcname)
	}
	return fmt.Sprintf("%v: %v(%v)", testname, funcname, argname)
}

func typeGuardExclude(needle Any, notAllowed []types.Any) bool {
	return !Contains(needle, notAllowed)
}

func tTypeError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v(%T), want %v(%T)", name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func tError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v, want %v", name, limitTestResultLength(got), limitTestResultLength(want))
}
func tTypeRun(t *testing.T, name string, got, want Any) {
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if got != want {
				if !reflect.DeepEqual(got, want) {
					tTypeError(t, name, got, want)
				}
			}
		})
	}
}

func tRun(t *testing.T, name string, got, want Any) {
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if got != want {
				if !reflect.DeepEqual(got, want) {
					tError(t, name, got, want)
				}
			}
		})
	}
}
