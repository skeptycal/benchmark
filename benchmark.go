// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

// Package benchmark contains utilities for macOS.
package benchmark

import (
	"fmt"
	"strings"
	"testing"

	"github.com/skeptycal/types"
)

const (
	// scaling factor (in powers of 2)
	defaultMaxScalingFactor = 6
	maxScalingFactor        = 10
)

type (
	AnyValue = types.AnyValue

	pooler interface {
		Get() *strings.Builder
		Release(sb *strings.Builder)
	}

	swimmer struct{ strings.Builder }
)

func (t swimmer) Get() *strings.Builder {
	return new(strings.Builder)
}

func (t swimmer) Release(sb *strings.Builder) {
	t.Reset()
}

func sbNonPool() pooler {
	return &swimmer{}
}

var (
// sb *strings.Builder = &strings.Builder{}
// NewPool                     = New()
// out        string = "" // global string return value
// global_n   int    = 0
// global_err error  = nil
// k          byte
)

type (
	Any = types.Any

	GetSetter interface {
		Get(key Any) (Any, error)
		Set(key Any, value Any) error
	}

	// Args implements a map of arguments
	Args interface {
		GetSetter
	}

	args map[Any]Any

	ArgSet []Args

	// Benchmark interface{}

	Benchmarks interface {
		Name() string
		Scale() int
		Count() int
		Next() Benchmark
		Setup() setupFunc
		Cleanup() cleanupFunc
	}

	benchmark2 struct {
		name    string   // name of benchmark test
		argSets []ArgSet // multiple runs with multiple argSets
		want    Any      // return value wanted
		wantErr bool     // is error wanted?
	}

	benchmarkSet2 struct {
		name          string      // name of benchmark set
		scale         int         // current scaling factor
		counter       int         // current trial counter
		runs          []benchmark // multiple runs if multiple Arg sets
		scalingFactor int         // max scaling factor for benchmark set (1-10)
		setup         setupFunc   // function used to setup benchmarks
		cleanup       cleanupFunc // function used to cleanup benchmarks
	}

	setupFunc   = func(set *benchmarkSet) error
	cleanupFunc = func(set *benchmarkSet) error
)

func NewBenchmarkSet2(name string, tests []benchmark, scalingFactor int, setup setupFunc, cleanup cleanupFunc) *benchmarkSet2 {
	if scalingFactor < 1 || scalingFactor > maxScalingFactor {
		scalingFactor = defaultMaxScalingFactor
	}

	if setup == nil {
		setup = defaultSetup
	}

	if cleanup == nil {
		cleanup = defaultCleanup
	}
	return &benchmarkSet2{
		name:          name,
		scale:         0,
		counter:       0,
		runs:          tests,
		scalingFactor: scalingFactor,
		setup:         setup,
		cleanup:       cleanup,
	}
}

func BenchmarkAll(b *testing.B) {
	benchmarks := benchmarkSet2{
		name:          "",
		scale:         0,
		counter:       0,
		runs:          []benchmark{},
		scalingFactor: defaultMaxScalingFactor,
		setup:         func(set *benchmarkSet) error { return nil },
		cleanup:       func(set *benchmarkSet) error { return nil },
	}

	for _, bb := range benchmarks.runs {
		for i := 0; i < b.N; i++ {
			fmt.Println(bb.name)
		}
	}

}

func defaultSetup(set *benchmarkSet) error {

	// set any global variables here
	// pass configuration options using default _config

	return nil
}

func defaultCleanup(set *benchmarkSet) error {
	return nil
}
