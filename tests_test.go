package benchmark

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

const wantSpam = false

const (
	maxArgChoices int     = 10
	one           uint64  = 1<<64 - 1
	halfBool      int64   = 1<<63 - 1
	halfHalfBool  int64   = 1<<62 - 1
	ratio         float64 = float64(one) / float64(halfBool)
	halfRatio     float64 = float64(halfBool) / float64(halfHalfBool)
)

func randBool() bool {
	return rand.Int63() >= halfHalfBool
}

func boolRatio(n int) float64 {
	m := boolSet(n)
	return float64(m[false]) / float64(m[true])
}

func boolSet(n int) map[bool]int {
	m := make(map[bool]int, 2)
	for i := 0; i < n; i++ {
		m[randBool()]++
	}
	return m
}

func boolSetGroup(n int) []float64 {
	list := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, boolRatio(n))
	}
	return list
}

func stdev(accepted float64, list []float64) float64 {
	w := len(list)
	sum := 0.0

	ac2 := accepted * accepted
	for i := 0; i < w; i++ {

		diff := math.Abs(ac2 - math.Pow(list[i], 2))
		sum += math.Sqrt(diff)
	}

	return sum / float64(w)
}

func TestRandBool(t *testing.T) {

	if wantSpam {

		for j := 0; j < 24; j++ {
			n := 1 << j
			ratio := boolRatio(n)
			// bs := boolSet(5)
			t.Errorf("randBool() n=%v ratio = %v, stdev = %v", n, ratio, nil)
		}
	}
}

func randomArg(n int) Any {
	if n < 0 || n > maxArgChoices {
		n = rand.Intn(10)
	}
	switch n {
	case 0: // nil
		return nil
	case 1: // int
		return rand.Intn(255)
	case 2: // uint64
		return rand.Uint64()
	case 3: // float
		return rand.Float64()
	case 4: // string
		return RandomString(rand.Intn(20) + 5)
	case 5: // []byte
		return []byte(RandomString(rand.Intn(20) + 5))
	case 6: // bool
		return randBool()
	case 7: // pointer
		var i int = 42
		return &i
	case 8: // slice
		return []Any{"4", "2"}
	case 9: // map
		return map[string]bool{"A": randomArg(6).(bool), "B": randomArg(6).(bool)}
	case 10: // bufio.Reader
		bufio.NewReader(os.Stdin)
	default:
	}
	return fmt.Sprintf("%v(%v)", "faker", n)
}

func randomTests(t *testing.T, name string, n int) []Tester {
	list := make([]Tester, 0, n)

	for i := 0; i < n; i++ {

		in := randomArg(6)

		nn := fmt.Sprintf("%v #%v(%v)", name, i, in)
		list = append(list, NewTest(t, nn, in, randomArg(-1), randomArg(-1), rand.Intn(1) == 1))
	}
	return list
}

// SampleTests generates n Tester set with n
// random sample Testers under the name given.
func SampleTests(t *testing.T, name string, n int) Tester {
	list := randomTests(t, name, n)
	return NewTestSet(t, name, list)
}

func TestSampleTests(t *testing.T) {
	if wantSpam {
		s := SampleTests(t, "sample", 100)
		s.Run()
	}
}

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
		TRun(t, tt.name, limitTestResultLength(tt.in), tt.want)
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

func TestNewTestSet(t *testing.T) {
	type args struct {
		t    *testing.T
		name string
		list []Tester
	}
	tests := []struct {
		name string
		args args
		want Tester
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTestSet(tt.args.t, tt.args.name, tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTestSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
