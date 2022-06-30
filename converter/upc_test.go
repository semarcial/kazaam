package converter

import (
	"strconv"
	"testing"

	"github.com/semarcial/kazaam/v5/registry"
)

func Test_Ucp_Convert(t *testing.T) {
	registry.RegisterConverter("upc", &Upc{})
	c := registry.GetConverter("upc")

	table := []struct {
		value     string
		arguments string
		expected  string
	}{
		{`"9526588315"`, `12`, `"009526588315"`},
		{`"0165719565"`, `12`, `"000165719565"`},
		{`"0000056787799"`, `12`, `"000056787799"`},
		{`"0000056787799"`, `8`, `"56787799"`},
		{`""`, `10`, `""`},
		{`"ACCF"`, `20`, `"ACCF"`},
	}

	for _, test := range table {
		v, e := c.Convert(getTestData(), []byte(test.value), []byte(strconv.Quote(test.arguments)))

		if e != nil {
			t.Error("error running convert function")
		}

		if string(v) != test.expected {
			t.Error("unexpected result from convert")
			t.Log("Expected: {}", test.expected)
			t.Log("Actual: {}", string(v))
		}
	}
}
