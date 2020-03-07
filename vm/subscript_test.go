package vm_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/antonmedv/expr"
)

type HelloFloat64 struct {
}

func (h HelloFloat64) Subscript(index string) float64 {
	return 124
}

type HelloString struct {
}

func (h HelloString) Subscript(index string) string {
	return "124"
}

func TestSubscriptStrValue(t *testing.T) {
	env := map[string]interface{}{
		"helloStr":     &HelloString{},
		"helloFloat64": &HelloFloat64{},
	}

	code := `helloStr["abc"] + "5"`

	program, err := expr.Compile(code, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	// fmt.Fprintln(os.Stdout, output)
	fmt.Println(output)

	if output != "1245" {
		t.Errorf("want 1235 got %v", output)
	}
}

func TestSubscriptFloat64Value(t *testing.T) {
	env := map[string]interface{}{
		"helloFloat64": &HelloFloat64{},
	}

	code := `helloFloat64["abc"] + 5`

	program, err := expr.Compile(code, expr.Env(env), expr.AsFloat64())
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	// fmt.Fprintln(os.Stdout, output)
	fmt.Println(output)

	if math.Abs(output.(float64)-129) > .00001 {
		t.Errorf("want 129 got %v", output)
	}
}
