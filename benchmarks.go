package benchmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/skeptycal/types"
)

var ValueOf = types.ValueOf

// NewBenchmarkSet returns a new set of Benchmark items.
func NewBenchmarkSet(b *testing.B, name string, set []Benchmark) BenchmarkSet {
	return &benchmarkSet{name: name, set: set}
}

type BmFunc = func(b *testing.B) []reflect.Value
type ReFunc = func(in []reflect.Value) []reflect.Value

func BenchmarkFunc(fn ReFunc, args []reflect.Value) BmFunc {
	if len(args) < 1 {
		return func(b *testing.B) []reflect.Value {
			return fn([]reflect.Value{ValueOf(b)})
		}
	}
	return func(b *testing.B) []reflect.Value {
		return fn(args)
	}
}

// NewBenchmark returns a new Benchmark item.
func NewBenchmark(name string, fn Any, args []Any) Benchmark {

	in := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}

	if v := reflect.ValueOf(fn); v.Kind() == reflect.Func {
		return &benchmark{name: name, fn: BenchmarkFunc(v.Call, in)}
	}
	return nil
}

type (
	Benchmark interface {
		Name() string
		Run(b *testing.B)
	}

	// BenchmarkSet is a collection of Benchmark items
	// that can be run with options applied.
	BenchmarkSet interface {
		Name() string
		Run(b *testing.B)
	}

	// benchmark implements Benchmark
	benchmark struct {
		name string
		fn   Any // func(b *testing.B)
		args []reflect.Value
	}

	// benchmarkSet implements BenchmarkSet
	benchmarkSet struct {
		name string
		set  []Benchmark
	}
)

func (bs *benchmarkSet) Run(b *testing.B) {
	for _, bb := range bs.set {
		name := fmt.Sprintf("%v - %v", bs.Name(), bb.Name())
		b.Run(name, func(b *testing.B) { bb.Run(b) })
	}
}
func (bm *benchmark) Run(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bm.fn.(BmFunc)(b)
	}
}
func (bs *benchmarkSet) Name() string { return bs.name }
func (bm *benchmark) Name() string    { return bm.name }
