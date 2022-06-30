package converter

import (
	"strconv"
	"testing"

	"github.com/semarcial/kazaam/v5/registry"
)

func TestBiglots_Ucp_Convert(t *testing.T) {
	registry.RegisterConverter("upc_biglots", &UpcBiglots{})
	c := registry.GetConverter("upc_biglots")

	table := []struct {
		value     string
		arguments string
		expected  string
	}{
		{`"9526588315"`, ``, `"009526588315"`},
		{`"0165719565"`, ``, `"000165719565"`},
		{`"0000056787799"`, ``, `"000056787799"`},
		{`""`, ``, `""`},
		{`"ACCF"`, ``, `"ACCF"`},
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
