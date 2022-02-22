package benchmark

// var defaultOptions = map[string]bool{
// 	"can_parallel":                            true,
// 	"use_global_return_value":                 true,
// 	"reallocate_local_variables_within_loops": true,
// }

// // NewTestSet returns a new set of Benchmark items.
// func NewTestSet(name string, set []Test) TestSet {
// 	return &testSet{&testOptions{name: name, options: defaultOptions}, set}
// }

// // NewTest returns a new Benchmark item.
// func NewTest(name string, fn func(t *testing.T)) Test {
// 	return &test{&testOptions{name: name, options: defaultOptions}, fn}
// }

// type (
// 	// Args []Any

// 	// Test is a Test itemthat can be run with
// 	// options applied.
// 	Test interface {
// 		Run(t *testing.T)
// 		Call(args ...Any) (string, error)

// 		TestOptioner
// 	}

// 	// TestSet is a collection of Test items
// 	// that can be run with options applied.
// 	TestSet interface {
// 		Run(t *testing.T)

// 		TestOptioner
// 	}

// 	// TestOptioner manages options for running tests.
// 	// The default options are:
// 	//	  defaultOptions = map[string]bool{
// 	//	 	"can_parallel":                            true,
// 	//	 	"use_global_return_value":                 true,
// 	//	 	"reallocate_local_variables_within_loops": true,
// 	//	 }
// 	TestOptioner interface {
// 		Name() string
// 		GetSetter
// 	}

// 	testOptions struct {
// 		name    string
// 		options map[string]bool
// 	}
// )

// func (to *testOptions) Name() string { return to.name }

// func (to *testOptions) Get(key Any) (Any, error) {
// 	if _, ok := key.(string); !ok {
// 		return nil, fmt.Errorf("key type must be string: %v(%T)", key, key)
// 	}
// 	if v, ok := to.options[key.(string)]; ok {
// 		return v, nil
// 	}
// 	return nil, fmt.Errorf("key not found: %v", key)
// }

// func (to *testOptions) Set(key Any, value Any) error {
// 	if _, ok := key.(string); !ok {
// 		return fmt.Errorf("key type must be string: %v(%T)", key, key)
// 	}
// 	if _, ok := value.(bool); !ok {
// 		return fmt.Errorf("value type must be bool: %v(%T)", value, value)
// 	}
// 	to.options[key.(string)] = value.(bool)
// 	return nil
// }

// // test implements Test
// type test struct {
// 	*testOptions
// 	fn func(t *testing.T)
// }

// func (tst *test) Run(t *testing.T)                 { tst.fn(t) }
// func (tst *test) Call(args ...Any) (string, error) { return "", nil }

// // testSet implements TestSet
// type testSet struct {
// 	*testOptions
// 	marks []Test
// }

// func (ts *testSet) Run(t *testing.T) {
// 	for _, bb := range ts.marks {
// 		name := fmt.Sprintf("%v - %v", ts.Name(), bb.Name())
// 		t.Run(name, func(t *testing.T) { bb.Run(t) })
// 	}
// }
