package converter

import (
	"strconv"
	"testing"

	"github.com/semarcial/kazaam/v5/registry"
)

func Test_Url_Convert(t *testing.T) {
	registry.RegisterConverter("url", &Url{})
	c := registry.GetConverter("url")

	table := []struct {
		value     string
		arguments string
		expected  string
	}{
		{`"//images.biglots.com/My+Arcade+Retro+Arcade+Machine+X?set=imageURL%5B%2"`, `https`, `"https://images.biglots.com/My+Arcade+Retro+Arcade+Machine+X?set=imageURL%5B%2"`},
		{`"//images.biglots.com/My+Arcade+Retro+Arcade+Machine+Xset=imageURL%5B%2"`, `http`, `"http://images.biglots.com/My+Arcade+Retro+Arcade+Machine+X?set=imageURL%5B%2"`},
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
