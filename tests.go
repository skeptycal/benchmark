package benchmark

import (
	"testing"
)

type (
	test struct {
		name    string
		in      Any
		got     Any
		want    Any
		wantErr bool
	}

	// Assert implements the Tester interface. It is
	// used for boolean only challenges. In addition
	// to working seamlessly with the standard library
	// testing package, it can return the bool
	// result for use in alternate data collection
	// or CI software.
	Assert interface {
		Tester
		Result() bool
	}

	// Random implements Tester and  creates a random
	// test that can be used to generate many varied
	// tests automatically.
	// After each use, Regenerate() can be called to
	// generate a new test.
	Random interface {
		Tester
		Regenerate()
	}

	// Custom implements Tester and can be used to
	// hook into existing software by passing in
	// the various test arguments with Hook().
	// Calling Hook() also calls Run() automaticaly.
	Custom interface {
		Tester
		Hook(name string, got, want Any, wantErr bool)
	}

	assert struct {
		name   string
		got    Any
		want   Any
		assert Assert
	}

	testSet struct {
		name string
		t    *testing.T
		list []test
	}

	// Tester implements an individual test. It may
	// be implemented by traditional tests,
	// asserts, random inputs, or custom code.
	Tester interface {

		// Run runs an individual test.
		Run(t *testing.T)
	}

	TestRunner interface {

		// Run runs all tests in the set.
		Run()
	}
)

func NewTestSet(t *testing.T, name string, list []test) TestRunner {
	return &testSet{
		t:    t,
		name: name,
		list: list,
	}
}

// Run runs all tests in the set.
func (ts *testSet) Run() {
	for _, tt := range ts.list {
		tt.Run(ts.t)
	}
}

// // Reset clears the list of tests
// func (ts *testSet) reset() {
// 	ts.list = []test{}
// }

// Run runs the individual test
func (tt *test) Run(t *testing.T) {
	tRunTest(t, tt)
}
