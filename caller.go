package benchmark

import "github.com/skeptycal/types"

var (

	// RetVal is the global Caller interface
	// used for the benchmarks of this package.
	RetVal = NewCaller(CallSetGlobalReturnValue)

	// globalReturnValue is the structure that holds
	// any global results from benchmarks.
	globalReturnValue interface{}
)

// CallSetGlobalReturnValue sets a global variable with the
// results of an operation. A function that implements
// this type is used in RetVal, the global Caller interface
// used for the benchmarks of this package.
//
// The input may be of any valid type. This is only useful
// during benchmark testing to eliminate possible compiler
// optimizations from the benchmark results, i.e.
//
// There are times when a good compiler may see a loop
// that seems to accomplish nothing useful or output any
// useful data and optimize it away. It may also be that
// the compiler could utilize different operations than
// were intended during a benchmark that would not be
// viable in a production environment. These types of
// events decrease the validity of benchmarks as profiling
// tools.
func CallSetGlobalReturnValue(any interface{}) []AnyValue {
	globalReturnValue = any
	return nil
}

type (

	// Caller is a generic function call interface.
	// The Call(interface{}) method calls the function
	// with any number of individual or variadic
	// arguments and returns a slice of AnyValue
	// return values.
	//
	//It can be enabled or disabled using the embedded
	// Enabler interface.
	//
	Caller interface {
		Call(v interface{}) []AnyValue
		types.Enabler
	}

	callerFunc func(v interface{}) []AnyValue

	caller struct {
		fn      callerFunc
		fnFalse callerFunc
		fnTrue  callerFunc
	}
)

// NewCaller returns a new instance of Caller
// with the function enabled.
func NewCaller(fn callerFunc) Caller {
	d := caller{
		fn:      fn,
		fnTrue:  fn,
		fnFalse: noop,
	}
	d.Enable()
	return &d
}

func (d *caller) Call(v interface{}) []AnyValue { return d.fn(v) }
func (d *caller) Enable()                       { d.fn = d.fnTrue }
func (d *caller) Disable()                      { d.fn = d.fnFalse }

func noop(any interface{}) []AnyValue { return nil } // noop function
