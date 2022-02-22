package benchmark

import (
	"reflect"
	"testing"

	"github.com/skeptycal/types"
)

var NewAnyValue = types.NewAnyValue

const (
	assertEqual     = "AssertEqual(%v): got %v, want %v"
	assertNotEqual  = "AssertNotEqual(%v): got %v, want %v"
	assertDeepEqual = "AssertDeepEqual(%v): got %v, want %v"
	assertSameType  = "AssertSameType(%v): got %v, want %v"
	assertSameKind  = "AssertSameKind(%v): got %v, want %v"
)

func CheckKind(got, want reflect.Kind) bool {
	return got == want
}

func AssertKind(t *testing.T, name string, got, want reflect.Kind) bool {
	if !CheckKind(got, want) {
		t.Errorf("kind is not equal(%v) = %v, want %v", name, got, want)
		return false
	}
	return true
}

// // Indirect will check if a is a pointer and return
// // the underlying value as needed. If a is not a pointer,
// // a is returned unchanged.
// func Indirect(a Any) Any {
// 	v := NewAnyValue(a)

// 	if v.Kind() != reflect.Ptr {
// 		return a
// 	}

// 	return v.TypeOf().Elem()
// }

func CompareFuncs(t *testing.T, name string, got, want Any) bool {

	g := NewAnyValue(got)

	// If pointer, run again with object
	if CheckKind(g.Kind(), reflect.Ptr) {
		return CompareFuncs(t, name, reflect.Indirect(g.ValueOf()), want)
	}

	// If pointer, run again with object
	w := NewAnyValue(want)
	if w.Kind() == reflect.Ptr {
		return CompareFuncs(t, name, got, reflect.Indirect(w.ValueOf()))
	}
	if g.Kind() != reflect.Func {
		t.Errorf("invalid type for function compare(%v) = %v, want %v", name, g.Kind(), reflect.Func)
		return false
	}
	return true
	// w := NewAnyValue(want)
}

func AssertEqual(t *testing.T, name string, got, want Any) bool {
	if got == want {
		return true
	}
	t.Errorf(assertEqual, name, got, want)
	return false
}

func AssertNotEqual(t *testing.T, name string, got, want Any) bool {
	if got != want {
		return true
	}
	t.Errorf(assertNotEqual, name, got, want)
	return false
}

func AssertDeepEqual(t *testing.T, name string, got, want Any) bool {
	if reflect.DeepEqual(got, want) {
		return true
	}
	t.Errorf(assertDeepEqual, name, got, want)
	return false
}

func AssertSameType(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got).TypeOf()
	w := NewAnyValue(want).TypeOf()

	if g == w {
		return true
	}
	t.Errorf(assertSameType, name, g, w)
	return false
}

func AssertSameKind(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got).Kind()
	w := NewAnyValue(want).Kind()

	if g == w {
		return true
	}
	t.Errorf(assertSameKind, name, g, w)
	return false
}
